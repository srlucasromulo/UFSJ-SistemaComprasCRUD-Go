package view

import (
	"fmt"
	"sistema-clp/model"
)

func MenuCrudVenda() int {

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Nova venda")
	fmt.Println("2 - Listar venda")
	fmt.Println("3 - Exclui venda")
	fmt.Println("4 - Mostrar total")
	fmt.Println("0 - Sair")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		return 0
	}

	return opt
}

func MenuListaVenda(v []model.Venda) {
	for i := range v {
		c := v[i].GetCliente()
		fmt.Println("Cliente: ", c.GetNome())
		fmt.Println("Data: ", v[i].GetData())
		fmt.Println("Numero: ", v[i].GetNumero())
		valor := PrintTotal(v)
		fmt.Println("Valor: ", valor)
	}
}

func PrintTotal(v []model.Venda) float32 {
	var num int
	var index int

	MenuListaVenda(v)

	fmt.Println("Digite o numero da venda")
	fmt.Scanf("%d", &num)

	for i := range v {
		if v[i].GetNumero() == num {
			index = i
		}
	}
	return model.ImprimeTotal(&v[index])
}
