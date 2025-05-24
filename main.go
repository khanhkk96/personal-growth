package main

import (
	"fmt"
	"log"
	configs "personal-growth/configs"
	injections "personal-growth/di"
	"personal-growth/docs"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @defaultModelRendering model
func main() {
	fmt.Print("Running service ...")
	loadConfig, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := configs.ConnectDB(&loadConfig)

	// db.AutoMigrate(
	// 	&entities.User{},
	// 	&entities.Project{},
	// 	&entities.Issue{},
	// 	&entities.Plan{},
	// 	&entities.Schedule{},
	// 	&entities.Article{},
	// 	&entities.Payment{},
	// )

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024, // limit 20 MB
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Static("/uploads", "./uploads")

	injections.MountAppRouters(app, db)

	// config swagger
	docs.SwaggerInfo.Title = "Swagger PersonalGrowth API"
	docs.SwaggerInfo.Description = "This is a PersonalGrowth API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// launch server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", loadConfig.ServerPort)))
}
