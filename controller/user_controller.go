package controller

import (
	"buat_docker/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service user.UserService
}

func NewUserController(service user.UserService) *userController {
	return &userController{service}
}

func (h *userController) CreateUser(c echo.Context) error {
	var input user.CreateUserInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	newUser, errUser := h.service.CreateUser(input)
	if errUser != nil {
		return c.JSON(http.StatusBadRequest, errUser)
	}

	return c.JSON(http.StatusCreated, newUser)
}

func (h *userController) GetAll(c echo.Context) error {
	users, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "data berhasil diambil",
		"data":    users,
	})
}
