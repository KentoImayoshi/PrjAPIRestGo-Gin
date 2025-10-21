package main

import (
	"github.com/kentoimayoshi/PrjAPIRestGo-Gin/database"
	"github.com/kentoimayoshi/PrjAPIRestGo-Gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
