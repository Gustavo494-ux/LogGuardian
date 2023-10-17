package models

import (
	enum "LogGuardian/src/enum/log"
	"errors"
	"time"
)

type Log struct {
	Id              uint64       `json:"id,omitempty"`
	CodigoErro      string       `json:"codigoErro,omitempty"`
	Tipo            enum.TipoLog `json:"tipo,omitempty"`
	NomePacote      string       `json:"nomePacote,omitempty"`
	NomeFuncao      string       `json:"nomeFuncao,omitempty"`
	Linha           int          `json:"linha,omitempty"`
	MensagemRetorno string       `json:"mensagemRetorno,omitempty"`
	MensagemErro    string       `json:"mensagemErro,omitempty"`
	DadosAdicionais interface{}  `json:"dadosAdicionais,omitempty"`
	DataHoraLog     time.Time    `json:"dataHoraLog,omitempty"`
}

func (log *Log) Validar() (err error) {
	if len(log.CodigoErro) == 0 {
		err = errors.New("o código é obrigatorio")
		return
	}

	if len(log.MensagemRetorno) == 0 || len(log.MensagemErro) == 0 {
		err = errors.New("a mensagem de retorno ou de erro deve ser forncecida")
		return
	}
	return
}
