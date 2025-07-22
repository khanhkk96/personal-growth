package main

import (
	"fmt"
	"log"
	configs "personal-growth/configs"
	injections "personal-growth/di"
	"personal-growth/docs"
	"personal-growth/handlers"
	"personal-growth/middlewares"
	"personal-growth/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @defaultModelRendering model
func main() {
	fmt.Print("Running service ...")
	loadedConfigs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := configs.ConnectDB(&loadedConfigs)

	// db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Project{},
	// 	&models.Issue{},
	// 	&models.Plan{},
	// 	&models.Schedule{},
	// 	&models.Article{},
	// 	&models.Payment{},
	// )

	// Seed admin acount if it doesn't exist
	handlers.InitAdmin(db)

	origins := viper.GetString("CORS_ORIGINS")
	corsConfig := cors.New(cors.Config{
		AllowOrigins: utils.Coalesce(&origins, "http://localhost:3000,http://localhost:5173"),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowCredentials: true,
	})
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024, // limit 20 MB

	})
	app.Use(corsConfig)
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
	log.Fatal(app.Listen(fmt.Sprintf(":%s", loadedConfigs.ServerPort)))
}
