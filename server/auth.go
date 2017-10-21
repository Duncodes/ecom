package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Duncodes/ecom/auth"
	"github.com/Duncodes/ecom/database"
)

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user database.UserCredential

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fulluser, err := user.VerifyUser()
	if err != nil {
		w.WriteHeader(400)

		response := map[string]interface{}{
			"error": "Wrong credentials",
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		log.Println("Wrong login credentilas")
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := auth.GenerateJWTTokken(user.Username, fulluser.UUID)
	if err != nil {
		// TODO : handle error
		log.Fatal(err)
	}

	response := map[string]interface{}{
		"token": token,
		"uuid":  fulluser.UUID,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// RegisterHandler this accepts a user detail of type
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	var err error
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		return
	}

	if err = user.CreateUser(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateJWTTokken(user.Username, user.UUID)

	if err != nil {
		log.Println(err)
		return
	}
	tokenresponse := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(tokenresponse)
}
