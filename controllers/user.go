package controllers

import (
	"fmt"
	"hashing-file/constants"
	"hashing-file/usecase/user"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserControllers struct {
	userService user.UserService
}

type UserControllersContract interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	User(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

func InitUserControllers(userService user.UserService) UserControllersContract {
	return &UserControllers{userService: userService}
}

func (u *UserControllers) Register(c *fiber.Ctx) (err error) {
	var data user.RegisterRequest
	if err = c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	res, err := u.userService.Register(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (u *UserControllers) Login(c *fiber.Ctx) (err error) {
	var data user.LoginRequest
	if err = c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	users, err := u.userService.Login(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    *users.Data.Token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Status(http.StatusOK).JSON(users)
}

func (u *UserControllers) User(c *fiber.Ctx) (err error) {
	cookie := c.Cookies("jwt")
	res, err := u.userService.User(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  constants.STATUS_UNAUTHORIZED,
			"Message": constants.MESSAGE_UNAUTHORIZED,
			"Data":    struct{}{},
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (u *UserControllers) Logout(c *fiber.Ctx) (err error) {
	data, err := u.userService.Logout()
	if err != nil {
		return
	}
	c.Cookie(data.Data.Cookies)
	res := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  data.Status,
		Message: fmt.Sprintf("%s Logout", data.Message),
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
