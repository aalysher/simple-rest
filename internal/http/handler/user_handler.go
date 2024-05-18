package handler

import (
	"fmt"
	"net/http"

	"github.com/aalysher/simple-rest/internal/models"
	"github.com/aalysher/simple-rest/internal/storage"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userRepo storage.UserRepository
}

func NewUserHandler(userRepo storage.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	user := new(models.User)
	fmt.Println("1")
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := u.userRepo.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}
