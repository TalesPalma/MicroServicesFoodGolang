package v1

import (
	"github.com/TalesPalma/PerdirComida/internal/handlers"
	"github.com/gorilla/mux"
)

func RoutersPedidos(router *mux.Router) {
	router.HandleFunc("/pedidos", handlers.RealizarPedito).Methods("POST")
	router.HandleFunc("/pedidos", handlers.VerificarPedidos).Methods("GET")
}
