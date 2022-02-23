package modelos

import (
	"errors"
	"strings"
)

//Essa estrutura representa a tabela de usuários no banco de dados
type Usuario struct {
	ID   uint64 `json: "id,omitempty"`
	Nome string `json: "nome,omitempty"`
	CPF  string `json: "cpf,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}
func (usuario *Usuario) validar() error {

	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.CPF == "" {
		return errors.New("o CPF é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.CPF = strings.TrimSpace(usuario.CPF)
}
