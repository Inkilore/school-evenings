package router

import (
	"backend/internal/database"
	"backend/internal/repo"
	"backend/internal/rule"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Api() {
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connection success")

	courseRepo := repo.NewCourseRepo(db)
	courseService := rule.NewCourseService(courseRepo)

	userRepo := repo.NewUserRepo(db)
	userService := rule.NewUserService(userRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	// app.Settings.BodyLimit = 10 * 1024 * 1024 // 10MB in bytes
	app.Server().MaxRequestBodySize = 10 * 1024 * 1024

	courseRouter(app, courseService)
	userRouter(app, userService)

	log.Fatal(app.Listen(":8080"))
}
