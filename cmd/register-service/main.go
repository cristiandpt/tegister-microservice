package main

import (
	"fmt"
	"github.com/cristiandpt/healthcare/register/internal/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	//"github.com/cristiandpt/healthcare/register/internal/model"
	"github.com/cristiandpt/healthcare/register/internal/middleware"
)

func main() {

	router := httprouter.New()
	router.POST("/api/register", middleware.LoggingMiddleware(handler.RegisterUser))
	router.GET("/password-recovery", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Println("Hello")
		fmt.Fprintf(w, "Hello")
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
