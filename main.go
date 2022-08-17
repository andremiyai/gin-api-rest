package main

import (
	"gin-api-rest/database"
	"gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// mock
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Gui Lima", CPF: "000.000.000-00", RG: "12.123.123-6"},
	// 	{Nome: "Gael Lima", CPF: "000.000.000-01", RG: "12.123.123-7"},
	// }
	routes.HandleRequests()
}
