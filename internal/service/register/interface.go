package register

import (
	"srm_arch/internal/entity"
)

type RegisterRepo interface {
	Login(req RegisterRequest) (entity.Users, error)
}
