package controller

import (
	"fmt"
	"sistema-clp/model"
	"sistema-clp/view"
)

func NovoProduto() model.Produto {

	codigo, nome, valor := view.MenuNovoProduto()
	var p model.Produto
	p.NovoProduto(codigo, nome, valor)
	return p
}

func AdicionarNovoProduto(p model.Produto, produtos []model.Produto) []model.Produto {

	produtos = append(produtos, p)
	return produtos
}

func AtualizaProduto(produtos []model.Produto) []model.Produto {

	view.MenuListaProduto(produtos)

	var cod int
	fmt.Println("Digite o codigo do produto: ")
	fmt.Scanf("%s", &cod)

	for i := range produtos {
		if produtos[i].GetCodigo() == cod {
			nome, valor := view.MenuAtualizaProduto()
			var p model.Produto
			p.NovoProduto(cod, nome, valor)
			produtos[i] = p
		}
	}
	return produtos
}

func ApagarProduto(produtos []model.Produto) []model.Produto {

	view.MenuListaProduto(produtos)

	var cod int
	fmt.Println("Digite o codigo do produto: ")
	fmt.Scanf("%d", &cod)

	var index int
	for i := range produtos {
		if produtos[i].GetCodigo() == cod {
			index = i
			break
		}
	}

	produtos = append(produtos[:index], produtos[index+1:]...)

	return produtos
}

func GetProduto(cod int, produtos []model.Produto) model.Produto {
	var index int
	for i := range produtos {
		if produtos[i].GetCodigo() == cod {
			index = i
		}
	}

	return produtos[index]
}
