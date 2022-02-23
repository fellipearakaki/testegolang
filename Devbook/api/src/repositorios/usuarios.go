package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Cria inside um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, cpf) values(?,?)",
	)
	if erro != nil {

		return 0, nil

	}
	defer statement.Close()

	//Faz a insercao
	resultado, erro := statement.Exec(usuario.Nome, usuario.CPF)
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

func (repositorio Usuarios) Buscar(nome string) ([]modelos.Usuario, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	linhas, erro := repositorio.db.Query(
		"select id, nome from usuarios where nome LIKE ?", nome)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscaPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select id, nome from usuarios where id = ?", ID)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, cpf = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.CPF, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id =?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorNome(nome string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, cpf from usuarios where nome =?", nome)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.CPF); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
