package main

import (
	"github.com/TalesPalma/PerdirComida/internal/db"
	serverconfig "github.com/TalesPalma/PerdirComida/internal/serverConfig"
)

const (
	port    = ":3000"
	AppName = "apiPedirComida"
)

func main() {
	db.InitDb()
	serverconfig.ServerSetupInit(port, AppName)
}
