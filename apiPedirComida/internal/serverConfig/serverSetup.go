package serverconfig

import (
	"log"
	"net/http"

	v1 "github.com/TalesPalma/PerdirComida/api/v1"
	"github.com/TalesPalma/PerdirComida/internal/etcd"
	"github.com/gorilla/mux"
)

func ServerSetupInit(port string, appName string) {
	r := mux.NewRouter()
	v1.RoutersPedidos(r)
	InitEtcd(appName, port)
	log.Println("Server runnit port", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Erro with init server:", err)
	}
}

func InitEtcd(appName string, port string) {
	clientEtcd := etcd.NewEtcdClient()
	err := clientEtcd.RegisterApi(appName, "3000")
	if err != nil {
		log.Fatal("Erro with register", err)
	}
	clientEtcd.DiscoveryService()
	clientEtcd.Close()
}
