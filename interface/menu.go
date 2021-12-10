package inter

import (
	"fmt"
)

func Menu() int {

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Menu clinte")
	fmt.Println("2 - Menu produto")
	fmt.Println("3 - Registrar venda")
	fmt.Println("0 - Sair")

	var opt int
	fmt.Scanf("%d", &opt)

	return opt
}

func MenuCRUD(tipo int) int {

	valor := func(tipo int) string {
		if tipo == 1 {
			return "cliente"
		}
		if tipo == 2 {
			return "produto"
		}
		return ""
	}(tipo)

	fmt.Println("Escolha uma opção:")
	fmt.Printf("1 - Novo %s\n", valor)
	fmt.Printf("1 - Lista %s\n", valor)
	fmt.Printf("1 - Atualiza %s\n", valor)
	fmt.Printf("1 - Exclui %s\n", valor)
	fmt.Println("0 - Sair")

	var opt int
	fmt.Scanf("%d", &opt)

	return opt
}

func MenuProduto() int {

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Novo clinte")
	fmt.Println("2 - Atualiza cliente")
	fmt.Println("3 - Excluir cliente")
	fmt.Println("4 - Listar cliente")
	fmt.Println("0 - Sair")

	return 0
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

func MenuInserirProduto() (int, string, float32) {

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

func RegistrarVenda() {
	//var numero int
	//var valor float32
}

//func Menu() {
func Menu2() {
	var opt int
	var nomePessoa string
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
