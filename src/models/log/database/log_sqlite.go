package models

import (
	models "LogGuardian/src/models/log"
	"LogGuardian/src/modules/GerenciadordeJson"
	"fmt"

	"github.com/jinzhu/copier"
)

type Log_sqlite struct {
	Id              uint64 `json:"id,omitempty" db:"id"`
	CodigoErro      string `json:"codigoErro,omitempty" db:"codigoErro"`
	Tipo            string `json:"tipo,omitempty" db:"tipo"`
	NomePacote      string `json:"nomePacote,omitempty" db:"nomePacote"`
	NomeFuncao      string `json:"nomeFuncao,omitempty" db:"nomeFuncao"`
	Linha           uint64 `json:"linha,omitempty" db:"linha"`
	MensagemRetorno string `json:"mensagemRetorno,omitempty" db:"mensagemRetorno"`
	MensagemErro    string `json:"mensagemErro,omitempty" db:"mensagemErro"`
	DadosAdicionais string `json:"dadosAdicionais,omitempty" db:"dadosAdicionais"`
	DataHoraLog     string `json:"dataHoraLog,omitempty" db:"dataHoraLog"`
}

// ImportarLog: recebe um models.Log e converte para um formato compatível com o models.Log_sqlite
func (log_sqlite *Log_sqlite) ImportarLog(log models.Log) {
	// Usando o Copier para copiar os campos idênticos
	if err := copier.Copy(&log_sqlite, log); err != nil {
		fmt.Println("Erro ao importar Log para Log_sqlite:", err)
	}
	log_sqlite.importarCamposPersonalizados(log)
}

// ExportarLog: converte o models.Log_sqlite para um formato compatível com o models.Log
func (log_sqlite *Log_sqlite) ExportarLog() (log models.Log) {
	// Usando o Copier para copiar os campos idênticos
	if err := copier.Copy(&log, log_sqlite); err != nil {
		fmt.Println("Erro ao exportar Log_sqlite para Log:", err)
	}
	return
}

// importarCamposPersonalizados: importa os campos que precisam algum tratamento
func (log_sqlite *Log_sqlite) importarCamposPersonalizados(log models.Log) {
	var err error

	log_sqlite.Tipo = string(log.Tipo)
	log_sqlite.DataHoraLog = log.DataHoraLog.Format("2006-01-02 15:04:05")

	log_sqlite.DadosAdicionais, err = GerenciadordeJson.InterfaceParaJsonString(log.DadosAdicionais)
	if err != nil {
		fmt.Println("erro ao converter uma interface para JsonString, err: ", err)
	}

}
