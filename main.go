package main

import (
	"LogGuardian/src/config"
	"LogGuardian/src/database"
	enum "LogGuardian/src/enum/log"
	interfaces "LogGuardian/src/interfaces/log"
	models "LogGuardian/src/models/log"
	"fmt"
	"math/rand"
	"time"
)

// Definindo um struct com três campos
type DadosAleatorios struct {
	Campo1 int
	Campo2 string
	Campo3 float64
}

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
	DadosAleatorios := DadosAleatorios{
		Campo1: rand.Intn(100),                                  // Número inteiro aleatório entre 0 e 99
		Campo2: fmt.Sprintf("TextoAleatorio%d", rand.Intn(100)), // Uma string com texto aleatório
		Campo3: rand.Float64() * 100,                            // Número de ponto flutuante aleatório entre 0 e 100
	}

	registro := models.Log{
		Id:              1,
		CodigoErro:      "ABC123",
		Tipo:            enum.TipoLog_Erro,
		NomePacote:      "Pacote1",
		NomeFuncao:      "Funcao1",
		Linha:           42,
		MensagemRetorno: "Retorno de erro",
		MensagemErro:    "Erro ocorrido",
		DadosAdicionais: DadosAleatorios,
		DataHoraLog:     time.Now(),
	}

	intt := interfaces.NovaInterfaceDeLog(&registro)
	intt.Criar()

	intt.BuscarPorTipo(string(enum.TipoLog_Erro))
}

func main() {
	TestarLog()
}
