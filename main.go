package main

import (
	"log"
	"money_management/config"
	"money_management/controllers"
	"money_management/database"
	"money_management/middlewares"
	"money_management/repositories"
	"money_management/services"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func init() {
	config.InitViper()
	utils.InitValidator()
	database.InitDB()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.CustomErrorHandler,
	})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	// user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// auth
	authService := services.NewAuthService(userService)
	authController := controllers.NewAuthController(authService)

	// transaction
	transactionRepository := repositories.NewTransactionRepository()
	transactionService := services.NewTransactionService(transactionRepository)
	transactionController := controllers.NewTransactionController(transactionService)

	// user
	app.Get("/users", userController.GetAllUsers)
	app.Get("/users/:id", userController.DetailUser)

	// auth
	app.Post("/login", authController.Login)
	app.Post("/register", authController.Register)
	app.Post("/logout", middlewares.AuthCheck, authController.Logout)

	// transaction
	app.Get("/transactions", middlewares.AuthCheck, transactionController.All)
	app.Get("/transactions/:id", middlewares.AuthCheck, transactionController.Detail)
	app.Delete("/transactions/:id", middlewares.AuthCheck, transactionController.Detail)
	app.Post("/transactions", middlewares.AuthCheck, transactionController.CreateTransaction)

	log.Fatal(app.Listen(":3000"))
}
