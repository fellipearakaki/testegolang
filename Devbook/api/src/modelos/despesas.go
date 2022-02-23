package modelos

import (
	"errors"
	"strings"
	"time"
)

type Despesas struct {
	ID        uint64
	Data      time.Time
	Id_forn   uint64
	Descricao string
	Valor     int
}

func (despesa *Despesas) Preparar() error {
	if erro := despesa.validar(); erro != nil {
		return erro
	}

	despesa.formatar()

	return nil
}

func (despesa *Despesas) validar() error {

	if despesa.Descricao == "" {
		return errors.New("A descrição é obrigatória e não pode estar em branco")
	}

	if despesa.Id_forn == 0 {
		return errors.New("o fornecedor é obrigatório e não pode estar em branco")
	}

	if despesa.Valor == 0 {
		return errors.New("O valor deve ser maior que zero e não pode estar em branco")
	}

	return nil
}

func (despesa *Despesas) formatar() {
	despesa.Descricao = strings.TrimSpace(despesa.Descricao)
	//despesa.Valor = strings.TrimSpace(despesa.Valor)
	//despesa.Data = strings.TrimSpace(despesa.Data)
}
