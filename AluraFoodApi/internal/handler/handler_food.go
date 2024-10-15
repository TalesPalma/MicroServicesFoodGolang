package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TalesPalma/AluraFood/internal/db"
	"github.com/TalesPalma/AluraFood/internal/models"
)

func GetFoods(w http.ResponseWriter, r *http.Request) {

	// Find all foods
	var foods []models.Food
	db.DB.Find(&foods)
	log.Println("Find all foods")

	// Send foods as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(foods)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func CreateFood(w http.ResponseWriter, r *http.Request) {

	log.Println("Create food")
	var food models.Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Create food
	log.Println(food)
	db.DB.Create(&food)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send food as response
	err := json.NewEncoder(w).Encode(food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
