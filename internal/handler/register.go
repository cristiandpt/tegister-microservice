package handler

import (
	"log"
	"io"
	"github.com/julienschmidt/httprouter"
	"net/http"
	jsonparse "encoding/json"
	"github.com/cristiandpt/healthcare/register/internal/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user model.UserRegisterDTO	
	postBody, readerErr := io.ReadAll(r.Body)
	if readerErr != nil {
		log.Println("Throw error")
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return 
	}
	err:= jsonparse.Unmarshal(postBody, &user)
	if err != nil {
		log.Printf("Throw error %s", err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"message": "User created successfully"}
	jsonResponse, _ := jsonparse.Marshal(response)
	w.Write(jsonResponse)
}
