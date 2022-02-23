package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Despesas representa um repositorio de fornecedores
type Despesas struct {
	db *sql.DB
}

//NovoRepositorioDeDespesas cria um repositorio de despesas
func NovoRepositorioDeDespesas(db *sql.DB) *Despesas {
	return &Despesas{db}
}

//Cria insere uma despesa no banco de dados
func (repositorio Despesas) Criar(despesa modelos.Despesas) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into despesas (data, descricao, id_forn, valor ) values(?,?,?,?)",
	)
	if erro != nil {

		return 0, nil

	}
	defer statement.Close()

	//Faz a insercao
	resultado, erro := statement.Exec(despesa.Data, despesa.Descricao, despesa.Id_forn, despesa.Valor)
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

func (repositorio Despesas) Buscar(nome string) ([]modelos.Despesas, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	linhas, erro := repositorio.db.Query(
		"select id, nome from fornecedor where nome LIKE ?", nome)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var despesas []modelos.Despesas

	for linhas.Next() {
		var despesa modelos.Despesas

		if erro = linhas.Scan(
			&despesa.ID,
			&despesa.Descricao,
		); erro != nil {
			return nil, erro
		}

		despesas = append(despesas, despesa)
	}

	return despesas, nil
}

func (repositorio Despesas) BuscaPorID(ID uint64) (modelos.Despesas, error) {
	linhas, erro := repositorio.db.Query("select id, descricao, valor from despesas where id = ?", ID)

	if erro != nil {
		return modelos.Despesas{}, erro
	}

	defer linhas.Close()

	var despesa modelos.Despesas

	if linhas.Next() {
		if erro = linhas.Scan(
			&despesa.ID,
			&despesa.Descricao,
			&despesa.Valor,
		); erro != nil {
			return modelos.Despesas{}, erro
		}
	}

	return despesa, nil
}

func (repositorio Despesas) Atualizar(ID uint64, despesa modelos.Despesas) error {
	statement, erro := repositorio.db.Prepare("update despesas set data = ?, descricao = ?, valor = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(despesa.Data, despesa.Descricao, despesa.Valor, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Despesas) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from despesas where id =?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
