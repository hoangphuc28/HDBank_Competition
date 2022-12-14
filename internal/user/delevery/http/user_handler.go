package userhttp

import (
	"context"
	usermodel "github.com/Zhoangp/HDBank_Competition/internal/user/model"
	otherapi "github.com/Zhoangp/HDBank_Competition/otherApi"
	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Register(ctx context.Context, user *usermodel.User) error
}
type userHandler struct {
	userUc UserUsecase
}

func NewUserHandler(userUc UserUsecase) *userHandler {
	return &userHandler{userUc}
}
func (hdl *userHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.User
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := hdl.userUc.Register(c.Request.Context(), &data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		s := otherapi.Register(data)
		c.JSON(200, gin.H{"result": s})
	}
}
