package register

import (
	"context"
	"github.com/uptrace/bun"
	"srm_arch/internal/entity"
	"srm_arch/internal/service/register"
)

type registerRepo struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *registerRepo {
	return &registerRepo{
		DB,
	}
}

func (r *registerRepo) Login(req register.RegisterRequest) (entity.Users, error) {
	user := new(entity.Users)
	err := r.NewSelect().Model(user).Where("phone_number = ? and password = ? and deleted_at IS NULL", req.PhoneNumber, req.Password).Scan(context.Background())
	if err != nil {
		return entity.Users{}, err
	}
	return *user, nil
}
