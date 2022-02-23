package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Fornecedores representa um repositorio de fornecedores
type Fornecedores struct {
	db *sql.DB
}

//NovoRepositorioDeFornecedoress cria um repositorio de fornecedores
func NovoRepositorioDeFornecedores(db *sql.DB) *Fornecedores {
	return &Fornecedores{db}
}

//Cria insere um fornecedor no banco de dados
func (repositorio Fornecedores) Criar(fornecedor modelos.Fornecedor) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into fornecedores (nome, cnpj) values(?,?)",
	)
	if erro != nil {

		return 0, nil

	}
	defer statement.Close()

	//Faz a insercao
	resultado, erro := statement.Exec(fornecedor.Nome, fornecedor.CNPJ)
	if erro != nil {
		return 0, nil
	}

	//Pega o valor do ultimo id inserido
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}
	return uint64(ultimoIDInserido), nil // precisa converter para uint pq LastInsertId() retorna um int
}

func (repositorio Fornecedores) Buscar(nome string) ([]modelos.Fornecedor, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	linhas, erro := repositorio.db.Query(
		"select id, nome from fornecedor where nome LIKE ?", nome)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var fornecedores []modelos.Fornecedor

	for linhas.Next() {
		var fornecedor modelos.Fornecedor

		if erro = linhas.Scan(
			&fornecedor.ID,
			&fornecedor.Nome,
		); erro != nil {
			return nil, erro
		}

		fornecedores = append(fornecedores, fornecedor)
	}

	return fornecedores, nil
}

func (repositorio Fornecedores) BuscaPorID(ID uint64) (modelos.Fornecedor, error) {
	linhas, erro := repositorio.db.Query("select id, nome from fornecedor where id = ?", ID)

	if erro != nil {
		return modelos.Fornecedor{}, erro
	}

	defer linhas.Close()

	var forn modelos.Fornecedor

	if linhas.Next() {
		if erro = linhas.Scan(
			&forn.ID,
			&forn.Nome,
		); erro != nil {
			return modelos.Fornecedor{}, erro
		}
	}

	return forn, nil
}

func (repositorio Fornecedores) Atualizar(ID uint64, fornecedor modelos.Fornecedor) error {
	statement, erro := repositorio.db.Prepare("update fornecedor set nome = ?, cnpj = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(fornecedor.Nome, fornecedor.CNPJ, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Fornecedores) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from fornecedor where id =?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
