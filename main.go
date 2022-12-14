package main

import (
	"fmt"
	"github.com/Zhoangp/HDBank_Competition/config"
	userhttp "github.com/Zhoangp/HDBank_Competition/internal/user/delevery/http"
	userrepository "github.com/Zhoangp/HDBank_Competition/internal/user/repository"
	userusecase "github.com/Zhoangp/HDBank_Competition/internal/user/usecase"
	otherapi "github.com/Zhoangp/HDBank_Competition/otherApi"
	"github.com/Zhoangp/HDBank_Competition/pkg/db/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

// "log"
// "os"
// "github.com/Zhoangp/HDBank_Competition/config"
// userrepository "github.com/Zhoangp/HDBank_Competition/internal/user/repository"
// "github.com/Zhoangp/HDBank_Competition/pkg/db/mysql"

func main() {
	env := os.Getenv("Env")
	fileName := "config/config-app.yml"
	if env == "app" {
		fmt.Println("app")
		fileName = "config/config-app.yml"
	}
	cfg, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("LoadConfig:", err)
		return
	}
	db, err := mysql.NewMySql(cfg)
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}
	userRepo := userrepository.NewUserRepository(db)
	userUc := userusecase.NewUserUseCase(userRepo)
	userHandler := userhttp.NewUserHandler(userUc)
	fmt.Println(otherapi.GetPublicKey())
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/register", userHandler.Register())
	r.Run(":" + cfg.App.Port)

	//otherapi.Register(usermodel.User{"zhoangp", "Abc12345!", "Nguyen Van A", "vana123@gmail.com", "0948407150", "777123777"})
}
