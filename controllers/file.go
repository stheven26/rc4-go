package controllers

import (
	"hashing-file/constants"
	"hashing-file/usecase/file"
	"hashing-file/usecase/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type fileControllers struct {
	fileService file.FileService
	userService user.UserService
}

type FileControllersContract interface {
	UploadDocument(c *fiber.Ctx) error
	EncryptDocument(c *fiber.Ctx) error
	DecryptDocument(c *fiber.Ctx) error
}

func InitFileControllers(file file.FileService, user user.UserService) FileControllersContract {
	return &fileControllers{
		fileService: file,
		userService: user,
	}
}

func (f *fileControllers) UploadDocument(c *fiber.Ctx) (err error) {
	cookie := c.Cookies("jwt")
	_, err = f.userService.User(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  constants.STATUS_UNAUTHORIZED,
			"Message": constants.MESSAGE_UNAUTHORIZED,
			"Data":    struct{}{},
		})
	}
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	res, err := f.fileService.UploadFile(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (f *fileControllers) EncryptDocument(c *fiber.Ctx) (err error) {
	// id := c.Params("id")
	cookie := c.Cookies("jwt")
	_, err = f.userService.User(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  constants.STATUS_UNAUTHORIZED,
			"Message": constants.MESSAGE_UNAUTHORIZED,
			"Data":    struct{}{},
		})
	}
	if constants.Key == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	res, err := f.fileService.EncryptFile(constants.ID, constants.Key, constants.Passphrase, constants.Data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	return c.Status(http.StatusOK).JSON(res)
}

func (f *fileControllers) DecryptDocument(c *fiber.Ctx) (err error) {
	// id := c.Params("id")
	cookie := c.Cookies("jwt")
	_, err = f.userService.User(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  constants.STATUS_UNAUTHORIZED,
			"Message": constants.MESSAGE_UNAUTHORIZED,
			"Data":    struct{}{},
		})
	}
	if constants.Key == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	res, err := f.fileService.DecryptFile(constants.ID, constants.Key, constants.Passphrase)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	return c.Status(http.StatusOK).JSON(res)
}
