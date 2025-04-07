package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"		
	"github.com/cristiandpt/healthcare/register/internal/handler" 
        //"github.com/cristiandpt/healthcare/register/internal/model"   
)

func main() {

	router := httprouter.New()
	router.POST("/api/register", handler.RegisterUser)
	router.GET("/password-recovery", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Println("Hello")
		fmt.Fprintf(w, "Hello")
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}



