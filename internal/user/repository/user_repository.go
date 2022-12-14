package userrepository

import (
	"context"

	usermodel "github.com/Zhoangp/HDBank_Competition/internal/user/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (r *userRepository) Create(ctx context.Context, user usermodel.UserDb) error {
	db := r.db.Begin()
	if err := db.Table(user.TableName()).Create(&user).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
