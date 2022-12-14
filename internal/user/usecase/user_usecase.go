package userusecase

import (
	"context"
	usermodel "github.com/Zhoangp/HDBank_Competition/internal/user/model"
)

type UserRepository interface {
	Create(context.Context, usermodel.UserDb) error
}
type userUseCase struct {
	userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *userUseCase {
	return &userUseCase{userRepository}
}
func (u *userUseCase) Register(ctx context.Context, data *usermodel.User) error {
	userDb := usermodel.UserDb{data.UserName, data.Password}
	userDb.PrepareCreate()
	u.userRepository.Create(ctx, userDb)
	return nil
}
