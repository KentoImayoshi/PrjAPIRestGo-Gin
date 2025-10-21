package main

import (
	"github.com/kentoimayoshi/PrjAPIRestGo-Gin/models"
	"github.com/kentoimayoshi/PrjAPIRestGo-Gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "00000000000", RG: "440000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "450000000"},
	}
	routes.HandleRequests()
}
