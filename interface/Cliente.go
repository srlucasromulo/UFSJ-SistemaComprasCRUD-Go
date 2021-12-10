package inter

import (
	"fmt"
)

func MenuCrudCliente() int {

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Novo cliente")
	fmt.Println("2 - Lista cliente")
	fmt.Println("3 - Atualiza cliente")
	fmt.Println("4 - Exclui cliente")
	fmt.Println("0 - Sair")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		return 0
	}

	return opt
}

func MenuNovoCliente() (string, string, string, int, int, int) {

	var nome, endereco, rg string
	var dia, mes, ano int

	fmt.Println("Digite o nome:")
	fmt.Scanf("%s", nome)
	fmt.Println("Digite o endereço:")
	fmt.Scanf("%s", endereco)
	fmt.Println("Digite o RG:")
	fmt.Scanf("%s", rg)
	fmt.Println("Digite a data de nascimento:")
	fmt.Scanf("%d-%d-%d", &dia, &mes, &ano)

	return nome, endereco, rg, dia, mes, ano
}

func MenuListaCliente() {}

func MenuAtualizaCliente() {}

func MenuExcluiCliente() {}
