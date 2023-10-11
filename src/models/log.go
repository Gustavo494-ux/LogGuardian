package models

import (
	enum "LogGuardian/src/enum/log"
	"errors"
)

type Log struct {
	Id              uint64       `json:"id,omitempty" db:"id"`
	Codigo          string       `json:"codigo,omitempty" db:"codigo"`
	Tipo            enum.TipoLog `json:"tipo,omitempty" db:"tipo"`
	NomePacote      string       `json:"nome,omitempty" db:"nome"`
	NomeFuncao      string       `json:"nomeFuncao,omitempty" db:"nomeFuncao"`
	Linha           int          `json:"linha,omitempty" db:"linha"`
	MensagemRetorno string       `json:"mensagemRetorno,omitempty" db:"mensagemRetorno"`
	MensagemErro    string       `json:"mensagemErro,omitempty" db:"mensagemErro"`
	DadosAdicionais interface{}  `json:"dadosAdicionais,omitempty" db:"dadosAdicionais"`
}

func (log *Log) Validar() (err error) {
	if len(log.Codigo) == 0 {
		err = errors.New("o código é obrigatorio")
		return
	}

	if len(log.MensagemRetorno) == 0 || len(log.MensagemErro) == 0 {
		err = errors.New("a mensagem de retorno ou de erro deve ser forncecida")
		return
	}
	return
}
