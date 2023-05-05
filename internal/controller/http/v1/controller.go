package v1

import (
	"srm_arch/internal/usecase/admin"
	"srm_arch/internal/usecase/programmer"
	"srm_arch/internal/usecase/register"
)

type Controller struct {
	useCase struct {
		*admin.AdminUseCase
		*programmer.ProgrammerUseCase
		*register.RegisterUseCase
	}
}

func NewController(a *admin.AdminUseCase, p *programmer.ProgrammerUseCase, r *register.RegisterUseCase) *Controller {
	return &Controller{struct {
		*admin.AdminUseCase
		*programmer.ProgrammerUseCase
		*register.RegisterUseCase
	}{a, p, r}}
}
