package view

import (
	"fmt"
	"sistema-clp/model"
)

func MenuCrudVenda() int {

	fmt.Println("|------MENU VENDA------|")
	fmt.Println("1 - Nova venda")
	fmt.Println("2 - Listar venda")
	fmt.Println("3 - Exclui venda")
	fmt.Println("4 - Detalhar venda")
	fmt.Println("0 - Sair")
	fmt.Println("|----------------------|")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		return 0
	}

	return opt
}

func MenuListaVenda(v []model.Venda) {

	fmt.Println("|--------VENDAS--------|")
	for i := range v {
		c := v[i].GetCliente()
		fmt.Println("Numero: ", v[i].GetNumero())
		fmt.Println("Cliente(RG): ", c.GetRg())
		fmt.Println("Data: ", v[i].GetData())
		fmt.Println("Valor: ", v[i].Total())
	}
	fmt.Println("|----------------------|")
}
