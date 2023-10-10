package config

import (
	enum "LogGuardian/src/enum/database"
	routines "LogGuardian/src/routines/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConexao      = ""
	DriverBanco        = ""
	CaminhoBancoSQLite = ""
)

// InicializarConfigurações: Realiza a configuração de variáveis necessárias para o projeto funcionar.
func InicializarConfigurações() {
	carregarVariáveisDeAmbiente()
	ConfigurarConexaoBanco()
}

// carregarVariáveisDeAmbiente: Inicializa as variáveis de ambiente
func carregarVariáveisDeAmbiente() {
	err := godotenv.Load("LogGuardian.env")
	if err != nil {
		log.Fatal(err)
	}

	CaminhoBancoSQLite = os.Getenv("CaminhoBancoSQLite")
}

// ConfigurarConexaoBanco: Definea string de conexão de acordo com o banco a ser utilizado
func ConfigurarConexaoBanco() {
	switch DriverBanco {
	case string(enum.DriverBanco_SQLite):
		{
			routines.ConfigurarBancoSQLite3(CaminhoBancoSQLite, &StringConexao, &DriverBanco)
		}
	default:
		{
			routines.ConfigurarBancoSQLite3(CaminhoBancoSQLite, &StringConexao, &DriverBanco)
		}
	}
}
