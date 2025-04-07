package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/cristiandpt/healthcare/register/internal/database/config"
	entities "github.com/cristiandpt/healthcare/register/internal/database/entity"
	"github.com/cristiandpt/healthcare/register/internal/handler"
	"github.com/julienschmidt/httprouter"

	//"github.com/cristiandpt/healthcare/register/internal/model"
	"github.com/cristiandpt/healthcare/register/internal/middleware"
)

func main() {

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	err = database.AutoMigrate(&entities.User{})

	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
		return
	}

	router := httprouter.New()
	router.POST("/api/register", middleware.LoggingMiddleware(handler.RegisterUser))
	router.GET("/password-recovery", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Println("Hello")
		fmt.Fprintf(w, "Hello")
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
