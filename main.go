package main

import (
	"LogGuardian/src/config"
	"LogGuardian/src/database"
	enum "LogGuardian/src/enum/log"
	"LogGuardian/src/interfaces"
	"LogGuardian/src/models"
	"fmt"
	"time"
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

func TestarLog() {
	registro := models.Log{
		Id:              1,
		Codigo:          "ABC123",
		Tipo:            enum.TipoLog_FuncaoExecutada,
		NomePacote:      "Pacote1",
		NomeFuncao:      "Funcao1",
		Linha:           42,
		MensagemRetorno: "Retorno de erro",
		MensagemErro:    "Erro ocorrido",
		DadosAdicionais: map[string]interface{}{"chave": "valor"},
		DataHoraLog:     time.Now(),
	}

	intt := interfaces.NovaInterfaceDeLog(&registro)
	intt.Criar()
}

func main() {
	TestarLog()
}
