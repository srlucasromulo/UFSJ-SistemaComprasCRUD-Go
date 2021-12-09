package main

import (
	"fmt"
	inter "sistema-clp/interface"
	"sistema-clp/structs"
)

func main() {
	inter.Menu()

	var prod structs.Produto
	item := new(structs.ItemVenda)
	// var venda structs.Venda
	// var cli structs.Cliente

	// cli.NovoCliente("Felipe", "Rua da morte", "MG-222", )
	// venda.NovaVenda()
	prod.NovoProduto(0, "teste", 10.5)

	item.NovoItemVenda(prod, 10.5, 2)

	imprimeTotal(item)
}

func imprimeTotal(modulo structs.Totalizavel) {
	fmt.Println(modulo.Total())
}
