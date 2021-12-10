package main

import (
	"fmt"
	"sistema-clp/controller"
	"sistema-clp/model"
	"sistema-clp/view"
)

func main() {
	clientes := make([]model.Cliente, 0)
	produtos := make([]model.Produto, 0)
	vendas := make([]model.Venda, 0)

	for {

		opt := view.Menu()

		switch opt {

		case 1:
			for {
				crud := view.MenuCrudCliente()
				switch crud {
				case 1:
					var c model.Cliente
					c = controller.NovoCliente()
					clientes = controller.AdicionarNovoCliente(c, clientes)
				case 2:
					view.MenuListaCliente(clientes)
				case 3:
					clientes = controller.AtualizaCliente(clientes)
				case 4:
					clientes = controller.ApagarCliente(clientes)

				case 0:
				default:
					fmt.Println("Opção inválida!!")
				}
				if crud == 0 {
					break
				}
			}

		case 2:
			for {
				crud := view.MenuCrudProduto()
				switch crud {
				case 1:
					var p model.Produto
					p = controller.NovoProduto()
					produtos = controller.AdicionarNovoProduto(p, produtos)
				case 2:
					view.MenuListaProduto(produtos)
				case 3:
					produtos = controller.AtualizaProduto(produtos)
				case 4:
					produtos = controller.ApagarProduto(produtos)

				case 0:
				default:
					fmt.Println("Opção inválida!!")
				}
				if crud == 0 {
					break
				}
			}
		case 3:
			for {
				crud := view.MenuCrudVenda()
				switch crud {
				case 1:
					var v model.Venda
					v = controller.NovaVenda(clientes, produtos)
					vendas = controller.AdicionarNovaVenda(v, vendas)
				case 2:
					view.MenuListaVenda(vendas)
				case 3:
					vendas = controller.ApagarVenda(vendas)
				case 4:
					controller.DetalharVenda(vendas)

				case 0:
				default:
					fmt.Println("Opção inválida!!")
				}
				if crud == 0 {
					break
				}
			}

		case 0:
		default:
			fmt.Println("Opção inválida!!")
		}

		if opt == 0 {
			break
		}
	}
}
