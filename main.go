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
				itens := make([]model.ItemVenda, 0)
				switch crud {
				case 1:
					var v model.Venda
					var rg string
					var num int
					view.MenuListaCliente(clientes)
					fmt.Println("Digite um RG: ")
					fmt.Scanf("%s", &rg)
					c := controller.GetCliente(rg, clientes)
					fmt.Println("Digite o numero da compra")
					fmt.Scanf("%d", &num)
					for {
						var opt int
						var cod int
						var item model.ItemVenda
						fmt.Println("1 - Adicionar Produto")
						fmt.Println("0 - Fechar venda")
						fmt.Scanf("%d", &opt)
						if opt == 0 {
							break
						}

						view.MenuListaProduto(produtos)
						fmt.Print("Digite o codigo do produto: ")
						fmt.Scanf("%d", &cod)
						produto := controller.GetProduto(cod, produtos)
						fmt.Print("Digite a quantidade de produto: ")
						fmt.Scanf("%d", &cod)
						valor := produto.GetValor()
						item.NovoItemVenda(produto, valor, cod)
						itens = append(itens, item)

					}
					var dia int
					var mes int
					var ano int
					fmt.Print("Digite o dia, mes e ano da venda: ")
					fmt.Scanf("%d %d %d", &dia, &mes, &ano)
					var bday model.Bday
					bday.NovoBday(dia, mes, ano)
					v = controller.NovaVenda(c, itens, num, bday)
					vendas = controller.AdicionarNovaVenda(v, vendas)
				case 2:
					view.MenuListaVenda(vendas)
				case 3:
					vendas = controller.ApagarVenda(vendas)
				case 4:
					fmt.Println(view.PrintTotal(vendas))

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
