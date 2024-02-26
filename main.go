package main

import (
	database "github.com/jsoeirosmp/api-go-gin/databse"
	"github.com/jsoeirosmp/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Joao Soeiro", CPF: "45875760870", RG: "500605067"},
	// 	{Nome: "Ana", CPF: "12345678954", RG: "699584905"},
	// }
	routes.HandleRequests()

}
