package controller

import (
	"fmt"
	"sistema-clp/model"
	"sistema-clp/view"
)

func NovaVenda(clientes []model.Cliente, produtos []model.Produto) model.Venda {

	view.MenuListaCliente(clientes)

	var rg string
	fmt.Println("Selecione o RG do cliente: ")
	fmt.Scanf("%s", &rg)

	cliente := GetCliente(rg, clientes)

	var numero int
	fmt.Println("Digite o número da venda: ")
	fmt.Scanf("%d", &numero)

	var dia, mes, ano int
	fmt.Println("Digite a data da venda (DD/MM/AAAA): ")
	fmt.Scanf("%d/%d/%d", &dia, &mes, &ano)

	var data model.Data
	data.NovaData(dia, mes, ano)

	var opt int
	itens := make([]model.ItemVenda, 0)

	for {
		fmt.Println("|----------------------|")
		fmt.Println("1 - Adicionar Produto")
		fmt.Println("0 - Fechar venda")
		fmt.Println("|----------------------|")
		fmt.Scanf("%d", &opt)

		if opt == 0 {
			break
		}

		view.MenuListaProduto(produtos)

		var codigo, quantidade int

		fmt.Print("Digite o codigo do produto: ")
		fmt.Scanf("%d", &codigo)

		fmt.Print("Digite a quantidade do produto: ")
		fmt.Scanf("%d", &quantidade)

		produto := GetProduto(codigo, produtos)

		var item model.ItemVenda
		item.NovoItemVenda(produto, produto.GetValor(), quantidade)

		itens = append(itens, item)
	}

	var v model.Venda
	v.NovaVenda(numero, data, cliente, itens)

	return v
}

func AdicionarNovaVenda(v model.Venda, vendas []model.Venda) []model.Venda {
	vendas = append(vendas, v)
	return vendas
}

func ApagarVenda(vendas []model.Venda) []model.Venda {

	view.MenuListaVenda(vendas)

	var num int
	fmt.Println("Digite o número da venda: ")
	fmt.Scanf("%d", &num)

	var index int
	for i := range vendas {
		if vendas[i].GetNumero() == num {
			index = i
		}
	}

	vendas = append(vendas[:index], vendas[index+1:]...)

	return vendas
}

func DetalharVenda(vendas []model.Venda) {

	view.MenuListaVenda(vendas)

	var numero int
	fmt.Println("Digite o número da venda: ")
	fmt.Scanf("%d", &numero)

	for i := range vendas {

		if vendas[i].GetNumero() == numero {

			c := vendas[i].GetCliente()

			fmt.Printf("|-----VENDA %d-----|\n", numero)
			fmt.Println("Cliente(RG): ", c.GetRg())
			fmt.Println("Data: ", vendas[i].GetData())
			fmt.Println("Valor: ", vendas[i].Total())

			itens := vendas[i].GetItemVenda()

			fmt.Println("Produtos:")
			for i := range itens {

				p := itens[i].GetProduto()

				fmt.Println("\tProduto: ", p.GetNome())
				fmt.Println("\tQuantidade: ", itens[i].GetQuantidade())
				fmt.Println("\tValor un.: ", itens[i].GetValor())
				fmt.Println("\tTotal: ", itens[i].Total())
			}
			fmt.Println("|------------------|")

		}
	}
}
