package routines

import (
	enum "LogGuardian/src/enum/database"
	"LogGuardian/src/modules/GerenciadorDeArquivos"
	"log"
)

// ConfigurarBancoSQLite3: Realiza a configuração do banco SQLite3
func ConfigurarBancoSQLite3(CaminhoBancoSQLite string, StringConexao *string, Driver *string) {
	criarBancoSQLite3SeNaoExistir(CaminhoBancoSQLite)
	configurarStringConexaoEDrive(CaminhoBancoSQLite, StringConexao, Driver)
}

// criarBancoSQLite3SeNaoExistir: Verifica se o banco de dados existe, se não existir ele será criado
func criarBancoSQLite3SeNaoExistir(CaminhoBancoSQLite string) {
	if err := GerenciadorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(CaminhoBancoSQLite); err != nil {
		log.Fatal(
			"ocorreu um erro ao criar o banco de dados SQLite3, caminho do banco: ", CaminhoBancoSQLite,
			"	error:	", err)
	}
}

// configurarStringConexaoEDrive altera via ponteiro o valor das variaveis globais.
func configurarStringConexaoEDrive(CaminhoBancoSQLite string, StringConexao *string, Driver *string) {
	*StringConexao = CaminhoBancoSQLite
	*Driver = string(enum.DriverBanco_SQLite)
}
