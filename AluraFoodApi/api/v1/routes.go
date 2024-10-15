package v1

import (
	"github.com/TalesPalma/AluraFood/internal/handler"
	"github.com/gorilla/mux"
)

func RoutesFood(router *mux.Router) {
	router.HandleFunc(
		"/Food",
		handler.GetFoods,
	).Methods("GET")

	router.HandleFunc(
		"/Food",
		handler.CreateFood,
	).Methods("POST")

}
