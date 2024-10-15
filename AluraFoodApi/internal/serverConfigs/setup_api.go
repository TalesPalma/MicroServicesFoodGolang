package serverconfigs

import (
	"log"
	"net/http"

	v1 "github.com/TalesPalma/AluraFood/api/v1"
	"github.com/TalesPalma/AluraFood/internal/db"
	"github.com/gorilla/mux"
)

func Setup_api(port string, nameApi string) {
	db.InitDatabse()          // Using var DB *gorm.DB for operation with database
	router := mux.NewRouter() // Init server with mux router
	v1.RoutesFood(router)     // V1 routes for food Api
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, router)) // Start server
}
