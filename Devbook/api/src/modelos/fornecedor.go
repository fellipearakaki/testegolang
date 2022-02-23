package modelos

import (
	"errors"
	"strings"
)

type Fornecedor struct {
	ID   uint64 `json:id, omitempty`
	Nome string `json:nome, omitempty`
	CNPJ string `json:nome, omitempty`
}

func (fornecedor *Fornecedor) Preparar() error {
	if erro := fornecedor.validar(); erro != nil {
		return erro
	}

	fornecedor.formatar()

	return nil
}

func (fornecedor *Fornecedor) validar() error {

	if fornecedor.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if fornecedor.CNPJ == "" {
		return errors.New("o CNPJ é obrigatório e não pode estar em branco")
	}

	return nil
}

func (fornecedor *Fornecedor) formatar() {
	fornecedor.Nome = strings.TrimSpace(fornecedor.Nome)
	fornecedor.CNPJ = strings.TrimSpace(fornecedor.CNPJ)
}
