package inter

import (
	"fmt"
	"time"
)

func Menu() {
	var opt int;
	nomePessoa string;
	enderecoPessoa string
	rgCliente string
	dataCliente time.Time
	codigoProduto string
	nomeProduto string
	valorProduto float32
	numeroVenda int
	dataVenda time.Time
	

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
		fmt.Println("Digite o endereço:")
		fmt.Println("Digite o RG:")
		fmt.Println("Digite a data de nascimento:")
	case 2:
		fmt.Println("Ola", opt)
		fmt.Println("Digite o código do produto:")
		fmt.Println("Digite o nome do produto:")
		fmt.Println("Digite o valor do produto:")
	case 3:
		fmt.Println("Ola", opt)
		fmt.Println("Digite o número da Venda:")
		fmt.Println("Insira a data da venda:")
		fmt.Println("Insira o cliente:")

	}
}
