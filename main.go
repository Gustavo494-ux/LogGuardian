package main

import (
	"LogGuardian/src/config"
	"LogGuardian/src/database"
	"fmt"
	"time"
	// "log"
)

func init() {
	config.InicializarConfigurações()
	if err := database.TestarConexao(); err != nil {
		fmt.Println(
			"ocorreu um erro ao realizar o teste de conexão com o banco de dados, Driver: ", config.DriverBanco,
			" string de conexão: ", config.StringConexao,
			" 	erro: ", err)

		time.Sleep(time.Minute * 5)
	}
}

func main() {
}
