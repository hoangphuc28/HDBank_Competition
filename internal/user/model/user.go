package usermodel

import "github.com/Zhoangp/HDBank_Competition/pkg/utils"

type User struct {
	UserName       string `json:"userName" gorm:"useName"`
	Password       string `json:"password" gorm:"password"`
	FullName       string `json:"fullName" gorm:"fullName"`
	Email          string `json:"email" gorm:"email"`
	Phone          string `json:"phone" gorm:"phone"`
	IdentityNumber string `json:"identityNumber" gorm:"identifyNumber"`
	Key            string `json:"key"`
}

type UserDb struct {
	Username string `json:"userName" gorm:"user_name"`
	Password string `json:"password" gorm:"password"`
}

func (UserDb) TableName() string {
	return "listPassword"
}
func (u *UserDb) PrepareCreate() error {
	// Hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
