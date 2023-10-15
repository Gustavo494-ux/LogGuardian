package repository

import (
	models "LogGuardian/src/models/log/database"
	"LogGuardian/src/modules/GerenciadordeJson"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Log struct {
	db *sqlx.DB
}

// NovoRepositorioLog: Cria uma nova instância do repositorio de Log
func NovoRepositorioLog(db *sqlx.DB) *Log {
	return &Log{db}
}

func (repositorio Log) Criar(log models.Log_sqlite) (logId uint64, err error) {
	jsonString, _ := GerenciadordeJson.InterfaceParaJsonString(log.DadosAdicionais)
	statement, err := repositorio.db.Exec(
		` INSERT INTO Log (codigoErro, tipo, nomePacote, nomeFuncao, linha, 
			mensagemRetorno,mensagemErro,dadosAdicionais
			) 
		VALUES(?,?,?,?,?,?,?,?); `,
		log.Codigo,
		log.Tipo,
		log.NomePacote,
		log.NomeFuncao,
		log.Linha,
		log.MensagemRetorno,
		log.MensagemErro,
		jsonString,
		// log.DataHoraLog,
	)
	if err != nil {
		logId = 0
		if err == sql.ErrNoRows {
			err = errors.New("nenhum registro afetado, verifique os dados fornecidos")
			return
		} else {
			return
		}
	}

	id, err := statement.LastInsertId()
	if err != nil {
		logId = 0
		return
	}
	logId = uint64(id)
	return
}
