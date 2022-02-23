package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasFornecedores = []Rota{
	{
		Uri:                "/fornecedores",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarFornecedor,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/fornecedores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarFornecedores,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/fornecedores/{fornecedorId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarFornecedor,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/fornecedores/{fornecedorId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarFornecedor,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/fornecedores/{fornecedorId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarFornecedor,
		RequerAutenticacao: false,
	},
}
