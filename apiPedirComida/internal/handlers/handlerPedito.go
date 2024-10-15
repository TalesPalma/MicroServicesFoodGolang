package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalesPalma/PerdirComida/internal/db"
	"github.com/TalesPalma/PerdirComida/internal/models"
)

func RealizarPedito(w http.ResponseWriter, r *http.Request) {

	var p models.PedidosDto
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	newProduct := models.NovoPedido(p)
	db.DB.Create(&newProduct)
	fmt.Println(newProduct)
	w.WriteHeader(http.StatusCreated)
}

func VerificarPedidos(w http.ResponseWriter, r *http.Request) {

	var p []models.Pedidos

	result := db.DB.Find(&p)

	if result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println(p)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	productsDto := models.PedidosListToDto(p)
	err := json.NewEncoder(w).Encode(productsDto)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
