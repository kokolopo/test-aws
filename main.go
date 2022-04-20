package main

import (
	"buat_docker/config"
	"buat_docker/controller"
	"buat_docker/database"
	"buat_docker/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("Bissmillah...")
}

func main() {
	config := config.InitConfiguration()
	database.InitDatabase(config)

	userRepo := user.NewUserRepository(database.DB)
	userService := user.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	e := echo.New()

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "hay selamat datang",
		})
	})

	api := e.Group("api/v1")
	api.POST("/users", userController.CreateUser)
	api.GET("/users", userController.GetAll)

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))

}
