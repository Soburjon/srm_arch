package register

import (
	"srm_arch/internal/entity"
)

type RegisterService struct {
	repo RegisterRepo
}

func NewService(repo RegisterRepo) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

func (r *RegisterService) Login(req RegisterRequest) (entity.Users, error) {
	return r.repo.Login(req)
}
