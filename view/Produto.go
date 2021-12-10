package view

import (
	"fmt"
	"sistema-clp/model"
)

func MenuCrudProduto() int {

	fmt.Println("|-----MENU PRODUTO-----|")
	fmt.Println("1 - Novo produto")
	fmt.Println("2 - Lista produto")
	fmt.Println("3 - Atualiza produto")
	fmt.Println("4 - Exclui produto")
	fmt.Println("0 - Sair")
	fmt.Println("|----------------------|")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		return 0
	}

	return opt
}

func MenuNovoProduto() (int, string, float32) {

	var codigo int
	var nome string
	var valor float32

	fmt.Println("|-----NOVO PRODUTO-----|")
	fmt.Println("Digite o nome do produto:")
	fmt.Scanf("%s", &nome)
	fmt.Println("Digite o código do produto:")
	fmt.Scanf("%d", &codigo)
	fmt.Println("Digite o valor do produto:")
	fmt.Scanf("%f", &valor)
	fmt.Println("|----------------------|")

	return codigo, nome, valor
}

func MenuListaProduto(produtos []model.Produto) {

	fmt.Println("|-------PRODUTOS-------|")
	for i := range produtos {
		fmt.Println("Nome: ", produtos[i].GetNome())
		fmt.Println("Codigo: ", produtos[i].GetCodigo())
		fmt.Println("Valor: ", produtos[i].GetValor())
	}
	fmt.Println("|----------------------|")
}

func MenuAtualizaProduto() (string, float32) {
	var nome string
	var valor float32

	fmt.Println("Digite o nome do produto: ")
	fmt.Scanf("%s", &nome)

	fmt.Println("Digite o valor do produto: ")
	fmt.Scanf("%f", &valor)

	return nome, valor
}
