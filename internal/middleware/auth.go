package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"srm_arch/internal/controller/errors"
	configs "srm_arch/internal/pkg/config"
	"srm_arch/internal/pkg/jwt"
)

// JWTRoleAuthorizer is a structure for a Role Authorizer type
type JWTRoleAuthorizer struct {
	enforcer   *casbin.Enforcer
	SigningKey []byte
	//	logger     logger.Logger
}

// NewJWTRoleAuthorizer creates and returns new Role Authorizer
func NewJWTRoleAuthorizer(cfg *configs.Configuration) (*JWTRoleAuthorizer, error) {
	enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
	if err != nil {
		log.Fatal("could not initialize new enforcer:", err.Error())
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		SigningKey: []byte(cfg.JWTSecretKey),
	}, nil
}

// NewAuthorizer returns middleware function to be used by fiber app for authorization
func NewAuthorizer(jwtra *JWTRoleAuthorizer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("Authorization")

		claims, err := jwt.ExtractClaims(accessToken, jwtra.SigningKey)
		if err != nil {
			log.Println("could not extract claims:", err)
			return err
		}

		role := claims["role"]

		ok, err := jwtra.enforcer.Enforce(role, c.Path(), c.Method())
		if err != nil {
			log.Println("could not enforce:", err)
			return err
		}

		if !ok {
			err = c.SendStatus(http.StatusForbidden)
			if err != nil {
				return err
			}
			return c.JSON(errors.ErrorResponse{
				Code:    http.StatusForbidden,
				Message: errors.NotEnoughRights,
			})
		}

		return c.Next()
	}
}
