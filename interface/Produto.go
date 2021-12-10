package inter

import (
	"fmt"
)

func MenuCrudProduto() int {

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Novo produto")
	fmt.Println("2 - Lista produto")
	fmt.Println("3 - Atualiza produto")
	fmt.Println("4 - Exclui produto")
	fmt.Println("0 - Sair")

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

	fmt.Println("Digite o código do produto:")
	fmt.Scanf("%d", &codigo)
	fmt.Println("Digite o nome do produto:")
	fmt.Scanf("%s", nome)
	fmt.Println("Digite o valor do produto:")
	fmt.Scanf("%f", &valor)

	return codigo, nome, valor
}

func MenuListaProduto() {}

func MenuAtualizaProduto() {}

func MenuExcluiProduto() {}
