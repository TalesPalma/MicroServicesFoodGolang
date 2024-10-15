package main

import (
	serverconfigs "github.com/TalesPalma/AluraFood/internal/serverConfigs"
)

const (
	port    = ":3001"
	nameApi = "AluraFoodApi"
)

func main() {
	serverconfigs.Setup_api(port, nameApi)
}
