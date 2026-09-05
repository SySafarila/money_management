package main

import (
	"log"
	"money_management/config"
	"money_management/database"
	"money_management/middlewares"
	"money_management/modules"
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
	userModule := modules.NewUserModule()

	// auth
	authModule := modules.NewAuthModule(userModule.Service)

	// transaction
	transactionModule := modules.NewTransactionModule()

	// user
	app.Get("/users", userModule.Controller.GetAllUsers)
	app.Get("/users/:id", userModule.Controller.DetailUser)

	// auth
	app.Post("/login", authModule.Controller.Login)
	app.Post("/register", authModule.Controller.Register)
	app.Post("/logout", middlewares.AuthCheck, authModule.Controller.Logout)

	// transaction
	app.Use("/transactions", middlewares.AuthCheck)
	app.Get("/transactions", transactionModule.Controller.All)
	app.Post("/transactions", transactionModule.Controller.CreateTransaction)
	app.Get("/transactions/:id", transactionModule.Controller.Detail)
	app.Patch("/transactions/:id", transactionModule.Controller.UpdateTransaction)
	app.Delete("/transactions/:id", transactionModule.Controller.Delete)

	log.Fatal(app.Listen(":3000"))
}
