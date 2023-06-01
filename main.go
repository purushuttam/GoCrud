package main

import (
	"fmt"
	"golang-fiber-crud/config"
	"golang-fiber-crud/controller"
	"golang-fiber-crud/model"
	"golang-fiber-crud/repository"
	"golang-fiber-crud/router"
	"golang-fiber-crud/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Print("Run Service........")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load enviroment variables", err)
	}

	// database
	db := config.ConnectionDb(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	//init repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	//init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	//Init controller
	noteController := controller.NewNoteController(noteService)

	// Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
