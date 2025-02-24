package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/configure/database"
	"github.com/prapawit201/Test_Forviz/library/http"
	"github.com/prapawit201/Test_Forviz/library/repository"
	"github.com/prapawit201/Test_Forviz/library/service"
)

var (
	app = fiber.New()
)

func main() {

	//DB connection
	dbConnection, err := database.DBConection()
	if err != nil {
		panic(err)
	}

	//init repository
	userRepository := repository.NewLibraryRepository(dbConnection)

	//init service
	userService := service.NewLibraryService(userRepository)

	//init http
	httpHandler := http.NewHTTPHandler(userService)
	http.RegisterHandlers(app, httpHandler)

	log.Fatal(app.Listen(":3000"))
}
