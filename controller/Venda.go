package controller

import (
	"fmt"
	"sistema-clp/model"
	"sistema-clp/view"
)

func NovaVenda(c model.Cliente, itens []model.ItemVenda, num int, b model.Bday) model.Venda {
	var v model.Venda
	v.NovaVenda(num, b, c, itens)
	return v
}

func AdicionarNovaVenda(v model.Venda, vendas []model.Venda) []model.Venda {
	vendas = append(vendas, v)
	return vendas
}

func ApagarVenda(v []model.Venda) []model.Venda {
	var index int
	var num int

	view.MenuListaVenda(v)

	fmt.Println("Digite o numero da venda")
	fmt.Scanf("%d", &num)

	for i := range v {
		if v[i].GetNumero() == num {
			index = i
		}
	}

	v = append(v[:index], v[index+1:]...)

	return v
}
