package view

import (
	"fmt"
	"sistema-clp/model"
)

func MenuCrudCliente() int {

	fmt.Println("|-----MENU CLIENTE-----|")
	fmt.Println("1 - Novo cliente")
	fmt.Println("2 - Lista cliente")
	fmt.Println("3 - Atualiza cliente")
	fmt.Println("4 - Exclui cliente")
	fmt.Println("0 - Sair")
	fmt.Println("|----------------------|")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		return 0
	}

	return opt
}

func MenuNovoCliente() (string, string, string, model.Data) {

	var nome, endereco, rg string
	var dia, mes, ano int

	fmt.Println("|-----NOVO CLIENTE-----|")
	fmt.Println("Digite o nome:")
	fmt.Scanf("%s", &nome)
	fmt.Println("Digite o endereço:")
	fmt.Scanf("%s", &endereco)
	fmt.Println("Digite o RG:")
	fmt.Scanf("%s", &rg)
	fmt.Println("Digite a data de nascimento (DD/MM/AAAA)):")
	fmt.Scanf("%d/%d/%d", &dia, &mes, &ano)
	fmt.Println("|----------------------|")

	var bday model.Data

	bday.NovaData(dia, mes, ano)

	return nome, endereco, rg, bday
}

func MenuListaCliente(clientes []model.Cliente) {

	fmt.Println("|-------CLIENTES-------|")
	for i := range clientes {
		fmt.Println("Nome: ", clientes[i].GetNome())
		fmt.Println("Endereço: ", clientes[i].GetEndereco())
		fmt.Println("Rg: ", clientes[i].GetRg())
		fmt.Println("Data de nascimento: ", clientes[i].GetBday())
	}
	fmt.Println("|----------------------|")
}

func MenuAtualizaCliente() (string, string) {
	var nome string
	var endereco string

	fmt.Println("Digite o nome do cliente: ")
	fmt.Scanf("%s", &nome)

	fmt.Println("Digite o endereço do cliente: ")
	fmt.Scanf("%s", &endereco)

	return nome, endereco
}
