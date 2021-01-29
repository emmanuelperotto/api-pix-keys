package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"pixkeys/entities"
	"pixkeys/infra/middlewares"
	"pixkeys/usecases"
)

func CreatePixKey(w http.ResponseWriter, r *http.Request) {
	var pixKey entities.PixKey
	if err := json.NewDecoder(r.Body).Decode(&pixKey); err != nil {
		log.Println("[Request Body Error]", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pixKey.AccountID = middlewares.CurrentAccountID
	pixKey, err := usecases.CreatePixKey(pixKey)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pixKey); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}

