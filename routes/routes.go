package routes

import (
	"hashing-file/controllers"
	"hashing-file/db"
	"hashing-file/domain/repository"
	"hashing-file/usecase/file"
	"hashing-file/usecase/user"

	"github.com/gofiber/fiber/v2"
)

var (
	//db
	connection = db.SetupDB()
	//repository
	userRepository = repository.InitUserRepository(connection)
	fileRepository = repository.InitFileRepository(connection)
	//usecase
	userService = user.NewService(userRepository)
	fileService = file.NewService(fileRepository)
	//controllers
	userControllers = controllers.InitUserControllers(userService)
	fileControllers = controllers.InitFileControllers(fileService, userService)
)

func Routes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1User := v1.Group("/user")
	{
		v1User.Get("", userControllers.User)
		v1User.Get("/all", userControllers.GetAllUser)
		v1User.Post("/register", userControllers.Register)
		v1User.Post("/login", userControllers.Login)
		v1User.Post("/logout", userControllers.Logout)
	}
	v1File := v1.Group("/file")
	{
		v1File.Get("", fileControllers.GetAllDocument)
		v1File.Post("/upload", fileControllers.UploadDocument)
		v1File.Post("/encrypt", fileControllers.EncryptDocument)
		v1File.Post("/decrypt", fileControllers.DecryptDocument)
	}
}
