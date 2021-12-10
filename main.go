package main

import (
	"fmt"
	control "sistema-clp/controller"
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
)

func main() {
	clientes := make([]structs.Cliente, 0)
	produtos := make([]structs.Produto, 0)
	//var vendas []structs.Venda

	for {

		opt := inter.Menu()

		switch opt {

		case 1:
			crud := inter.MenuCrudCliente()
			switch crud {
			case 1:
				clientes = control.NovoCliente(clientes)
			case 2:
				inter.MenuListaCliente(clientes)
			case 3:
				var rg string
				inter.MenuListaCliente(clientes)
				fmt.Println("Digite o RG do cliente: ")
				fmt.Scanf("%s", &rg)
				control.AtualizaCliente(rg, clientes)
			case 4:
				var rg string
				inter.MenuListaCliente(clientes)
				fmt.Println("Digite o RG do cliente: ")
				fmt.Scanf("%s", &rg)
				clientes = control.ApagarCliente(rg, clientes)
			default:
				fmt.Println("Opção inválida!!")
			}

		case 2:
			crud := inter.MenuCrudProduto()
			switch crud {
			case 1:
				produtos = control.NovoProduto(produtos)
			case 2:
				inter.MenuListaProduto(produtos)
			case 3:
				var cod int
				inter.MenuListaProduto(produtos)
				fmt.Println("Digite o codido do produto: ")
				fmt.Scanf("%d", &cod)
				control.AtualizaProduto(cod, produtos)
			case 4:
				var cod int
				inter.MenuListaProduto(produtos)
				fmt.Println("Digite o codido do produto: ")
				fmt.Scanf("%d", &cod)
				produtos = control.ApagarProduto(cod, produtos)
			default:
				fmt.Println("Opção inválida!!")
			}

		default:
			fmt.Println("Opção inválida!!")
		}

		if opt == 0 {
			break
		}
	}
}

func imprimeTotal(modulo structs.Totalizavel) {
	fmt.Println(modulo.Total())
}
