package register

import (
	"srm_arch/internal/entity"
	"srm_arch/internal/service/register"
)

type RegisterUseCase struct {
	service RegisterService
}

func NewUseCase(s RegisterService) *RegisterUseCase {
	return &RegisterUseCase{
		service: s,
	}
}

func (r *RegisterUseCase) Login(req register.RegisterRequest) (entity.Users, error) {
	res, err := r.service.Login(req)
	if err != nil {
		return entity.Users{}, err
	}
	return res, nil
}
