package main

import (
	"fmt"
	"log"
	"personal-growth/config"
	"personal-growth/controller"
	"personal-growth/docs"
	"personal-growth/middlewares"
	"personal-growth/model"
	"personal-growth/repository"
	"personal-growth/router"
	"personal-growth/service"

	"github.com/go-playground/validator/v10"
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
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectDB(&loadConfig)
	validate := validator.New()
	db.AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Issue{},
		&model.Plan{},
		&model.Schedule{},
		&model.Article{},
	)

	// Auth
	userRepository := repository.NewBaseRepository[model.User](db)      //Init respository
	authService := service.NewAuthServiceImpl(userRepository, validate) //Init service
	authController := controller.NewAuthController(authService)         //Init controller
	authRoute := router.NewAuthRouter(authController, db)               //Init Routes

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024, // limit 20 MB
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Static("/uploads", "./uploads")

	app.Mount("/api", authRoute)

	docs.SwaggerInfo.Title = "Swagger PersonalGrowth API"
	docs.SwaggerInfo.Description = "This is a PersonalGrowth API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", loadConfig.ServerPort)))
}
