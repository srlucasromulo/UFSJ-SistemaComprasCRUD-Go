package inter

import (
	"fmt"
)

func Menu() {
	var opt int
	var nomePessoa string
	var enderecoPessoa string
	var rgCliente string
	var diaDataCliente int
	var mesDataCliente int
	var anoDataCliente int
	var codigoProduto int
	var nomeProduto string
	var valorProduto float32
	var numeroVenda int
	var diaDataVenda int
	var mesDataVenda int
	var anoDataVenda int

	fmt.Println("Seja bem vindo!")
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Inserir clinte")
	fmt.Println("2 - Inserir produto")
	fmt.Println("3 - Registrar venda")
	fmt.Println("4 - Sair")

	fmt.Scanf("%d", &opt)

	switch opt {
	case 1:
		fmt.Println("Ola", opt)
		fmt.Println("Digite o nome:")
		fmt.Scanf("%s", nomePessoa)
		fmt.Println("Digite o endereço:")
		fmt.Scanf("%s", enderecoPessoa)
		fmt.Println("Digite o RG:")
		fmt.Scanf("%s", rgCliente)
		fmt.Println("Digite a data de nascimento:")
		fmt.Scanf("%d-%d-%d", &diaDataCliente, &mesDataCliente, &anoDataCliente)
	case 2:
		fmt.Println("Ola", opt)
		fmt.Println("Digite o código do produto:")
		fmt.Scanf("%d", codigoProduto)
		fmt.Println("Digite o nome do produto:")
		fmt.Scanf("%s", nomeProduto)
		fmt.Println("Digite o valor do produto:")
		fmt.Scanf("%f", &valorProduto)
	case 3:
		fmt.Println("Ola", opt)
		fmt.Println("Digite o número da Venda:")
		fmt.Scanf("%d", &numeroVenda)
		fmt.Println("Insira a data da venda:")
		fmt.Scanf("%d-%d-%d", &diaDataVenda, &mesDataVenda, &anoDataVenda)
		fmt.Println("O cliente é:", nomePessoa) //listar clientes
	}
}
