package main

import (
	"apiRestGo/database"
	"apiRestGo/models"
	"apiRestGo/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Fernando Moreira", Historia: "Programador Go"},
		{Id: 2, Nome: "Fernando Moreira", Historia: "Programador Python"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando api rest em Go")
	routes.HandleRequest()
}
