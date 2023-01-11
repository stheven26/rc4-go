package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	LoginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	DefaultResponse struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	RegisterData struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
	}
	RegisterResponse struct {
		Status  string       `json:"status"`
		Message string       `json:"message"`
		Data    RegisterData `json:"data"`
	}
	LoginData struct {
		Email string  `json:"email"`
		Token *string `json:"token"`
	}
	LoginResponse struct {
		Status  string    `json:"status"`
		Message string    `json:"message"`
		Data    LoginData `json:"data"`
	}
	UserData struct {
		User    interface{} `json:"user"`
		Expires int64       `json:"expires"`
		Time    time.Time   `json:"time"`
	}
	UserResponse struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    UserData `json:"data"`
	}
	LogoutData struct {
		Cookies *fiber.Cookie `json:"token"`
	}
	LogoutResponse struct {
		Status  string     `json:"status"`
		Message string     `json:"message"`
		Data    LogoutData `json:"data"`
	}
)
