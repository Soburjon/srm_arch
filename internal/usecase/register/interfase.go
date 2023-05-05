package register

import (
	"srm_arch/internal/entity"
	"srm_arch/internal/service/register"
)

type RegisterService interface {
	Login(req register.RegisterRequest) (entity.Users, error)
}
