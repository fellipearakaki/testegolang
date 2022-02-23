package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasDespesas = []Rota{
	{
		Uri:                "/despesas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarDespesa,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/despesas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDespesas,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/despesas/{despesaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDespesa,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/despesas/{despesaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarDespesa,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/despesas/{despesaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarDespesa,
		RequerAutenticacao: false,
	},
}
