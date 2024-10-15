package serverconfig

import (
	"log"
	"net/http"

	v1 "github.com/TalesPalma/PerdirComida/api/v1"
	"github.com/gorilla/mux"
)

func ServerSetupInit(port string, appName string) {
	r := mux.NewRouter()
	v1.RoutersPedidos(r)
	log.Println("Server runnit port", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Erro with init server:", err)
	}
}
