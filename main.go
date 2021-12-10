package main

import (
	"fmt"
	control "sistema-clp/controller"
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
)

func main() {

	var clientes []structs.Cliente
	var produtos []structs.Produto
	//var vendas []structs.Venda

	for {

		opt := inter.Menu()

		switch opt {

		case 1:
			crud := inter.MenuCrudCliente()
			switch crud {
			case 1:
				control.NovoCliente(clientes)

			case 2:
				continue
			case 3:
				continue
			case 4:
				continue
			default:
				fmt.Println("Opção inválida!!")
			}

		case 2:
			crud := inter.MenuCrudProduto()
			switch crud {
			case 1:
				control.NovoProduto(produtos)

			case 2:
				continue
			case 3:
				continue
			case 4:
				continue
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
