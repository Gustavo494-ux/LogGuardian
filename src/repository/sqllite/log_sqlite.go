package repository

import (
	enum "LogGuardian/src/enum/log"
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

// Criar: salva um log no banco de dados
func (repositorio Log) Criar(log models.Log_sqlite) (logId uint64, err error) {
	jsonString, _ := GerenciadordeJson.InterfaceParaJsonString(log.DadosAdicionais)
	statement, err := repositorio.db.Exec(
		` INSERT INTO Log (codigoErro, tipo, nomePacote, nomeFuncao, linha, 
			mensagemRetorno,mensagemErro,dadosAdicionais,dataHoraLog
			) 
		VALUES(?,?,?,?,?,?,?,?,?); `,
		log.CodigoErro,
		log.Tipo,
		log.NomePacote,
		log.NomeFuncao,
		log.Linha,
		log.MensagemRetorno,
		log.MensagemErro,
		jsonString,
		log.DataHoraLog,
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

// BuscarPorTipo: Busca todos os logs do tipo informado
func (repositorio Log) BuscarPorTipo(tipoLog enum.TipoLog) (logBanco []models.Log_sqlite, err error) {
	err = repositorio.db.Select(&logBanco,
		`SELECT id, codigoErro, tipo, nomePacote, nomeFuncao, linha, mensagemRetorno, 
		mensagemErro, dadosAdicionais, dataHoraLog FROM Log WHERE tipo = ?`,
		string(tipoLog),
	)
	if err == sql.ErrNoRows {
		err = errors.New("nenhum usuário encontrado, verifique os dados fornecidos")
	}
	return
}
