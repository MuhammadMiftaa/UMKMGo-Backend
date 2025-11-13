package routes

import (
	"sapaUMKM-backend/config/redis"
	"sapaUMKM-backend/interface/http/handler"
	"sapaUMKM-backend/interface/http/middleware"
	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(version fiber.Router, db *gorm.DB, redis redis.RedisRepository) {
	User_repo := repository.NewUsersRepository(db)
	User_serv := service.NewUsersService(User_repo, redis)

	User_handler := handler.NewUsersHandler(User_serv)

	webAuth := version.Group("/webauth")
	{
		webAuth.Post("login", User_handler.Login)
		webAuth.Post("register", User_handler.Register)
		webAuth.Put("/profile", User_handler.UpdateProfile)
	}

	mobileAuth := version.Group("/mobileauth")
	{
		mobileAuth.Post("login", User_handler.Login)
		mobileAuth.Post("register", User_handler.Register)
		mobileAuth.Post("send/otp", User_handler.SendOTP)
		mobileAuth.Post("verify/otp", User_handler.VerifyOTP)
	}

	version.Use(middleware.AuthMiddleware(), middleware.ContextMiddleware())

	users := version.Group("/users")
	{
		users.Get("/", User_handler.GetAllUsers)
		users.Get(":id", User_handler.GetUserByID)
		users.Put(":id", User_handler.UpdateUser)
		users.Delete(":id", User_handler.DeleteUser)
	}

	version.Get("/permissions", User_handler.GetListPermissions)
	version.Get("/role-permissions", User_handler.GetListRolePermissions)
	version.Post("/role-permissions", User_handler.UpdateRolePermissions)
}
