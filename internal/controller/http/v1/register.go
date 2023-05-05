package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"srm_arch/internal/pkg/utils"
	"srm_arch/internal/service/register"
	"strings"
)

// Login method user id va authorize qaytaradi
// @Description user id va authorize qaytaradi
// @Summary user id va authorize qaytaradi
// @Tags register
// @Accept json
// @Produce json
// @Param login body models.RegisterRequest true "login"
// @Success 200 {object} models.LoginResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /register/login/ [POST]
func (a *Controller) Login(c *fiber.Ctx) error {
	req := register.RegisterRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !strings.Contains(req.PhoneNumber, "+") ||
		len(req.PhoneNumber) != 13 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "phone number is not correctly filled",
		})
	}
	res, err := a.useCase.Login(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println(res)
	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(res.ID, map[string]string{
		"role": res.Role,
	})
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(register.LoginResponse{
		UserID:    res.ID,
		Authorize: tokens.Access,
	})
}
