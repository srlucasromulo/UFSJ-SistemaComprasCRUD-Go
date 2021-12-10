package controller

import (
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
)

func NovoProduto(produtos []structs.Produto) {

	codigo, nome, valor := inter.MenuNovoProduto()
	var p structs.Produto
	p.NovoProduto(codigo, nome, valor)
	// adicionar em produtos
}
