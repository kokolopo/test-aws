package main

import (
	"buat_docker/controller"
	"buat_docker/database"
	"buat_docker/user"

	"github.com/labstack/echo/v4"
)

// func init() {
// 	fmt.Println("Bissmillah...")
// }

func main() {
	database.InitDatabase()

	userRepo := user.NewUserRepository(database.DB)
	userService := user.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	e := echo.New()

	api := e.Group("api/v1")
	api.POST("/users", userController.CreateUser)
	api.GET("/users", userController.GetAll)

	e.Logger.Fatal(e.Start(":8080"))

}
