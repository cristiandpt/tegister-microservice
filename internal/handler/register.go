package handler

import (
	jsonparse "encoding/json"
	database "github.com/cristiandpt/healthcare/register/internal/database/config"
	entities "github.com/cristiandpt/healthcare/register/internal/database/entity"
	"github.com/cristiandpt/healthcare/register/internal/model"
	"github.com/cristiandpt/healthcare/register/util"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db := database.GetDB()

	var user model.UserRegisterDTO
	postBody, readerErr := io.ReadAll(r.Body)
	if readerErr != nil {
		log.Println("Throw error")
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err := jsonparse.Unmarshal(postBody, &user)
	if err != nil {
		log.Printf("Throw error %s", err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	encryptedPassword, err := util.EncryptPassword(user.Password)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	newUser := entities.User{
		Username:     user.Email,
		Email:        user.Email,
		PasswordHash: encryptedPassword,
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
		return
	}

	log.Printf("GORM with pgx User registered successfully with ID: %s, Username: %s, Email: %s", newUser.ID, newUser.Username, newUser.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]any{"message": "User created successfully"}
	jsonResponse, _ := jsonparse.Marshal(response)
	w.Write(jsonResponse)
}
