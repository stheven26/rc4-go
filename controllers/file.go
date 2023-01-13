package controllers

import (
	"fmt"
	"hashing-file/constants"
	"hashing-file/usecase/file"
	"hashing-file/usecase/user"
	"net/http"
	"os"
	"time"

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
	if constants.EncryptID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
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
	_, err = f.fileService.EncryptFile(constants.EncryptID, constants.Key, constants.Passphrase, constants.Data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	constants.EncryptID = ""
	return c.Status(http.StatusOK).Download(fmt.Sprintf("./%s", constants.Key), fmt.Sprintf("encrypt-%v", time.Now().UnixMicro()))
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
	if constants.DecryptID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
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
	_, err = f.fileService.DecryptFile(constants.DecryptID, constants.Key, constants.Passphrase)
	if err != nil {
		fmt.Println("error disini")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  constants.STATUS_FAILED,
			"Message": constants.MESSAGE_FAILED,
			"Data":    struct{}{},
		})
	}
	constants.DecryptID = ""
	defer os.Remove(fmt.Sprintf("./%s", constants.Key))
	defer os.Remove(fmt.Sprintf("./decrypt-%s", constants.Key))
	return c.Status(http.StatusOK).Download(fmt.Sprintf("./decrypt-%s", constants.Key), fmt.Sprintf("decrypt-%v", time.Now().UnixMicro()))
}
