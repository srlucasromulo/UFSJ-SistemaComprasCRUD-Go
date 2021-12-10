package controller

import (
	"fmt"
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
)

func NovoProduto(produtos []structs.Produto) []structs.Produto {

	codigo, nome, valor := inter.MenuNovoProduto()
	var p structs.Produto
	p.NovoProduto(codigo, nome, valor)
	produtos = append(produtos, p)
	fmt.Println(produtos[0].GetNome())

	return produtos
}

func AtualizaProduto(cod int, produtos []structs.Produto) {
	for i := range produtos {
		if produtos[i].GetCodigo() == cod {
			nome, valor := inter.MenuAtualizaProduto()
			produtos[i].SetValor(valor)
			produtos[i].SetNome(nome)
		}
	}
}

func ApagarProduto(cod int, produtos []structs.Produto) []structs.Produto {
	var index int
	for i := range produtos {
		if produtos[i].GetCodigo() == cod {
			index = i
		}
	}

	produtos = append(produtos[:index], produtos[index+1:]...)

	return produtos
}
