package main

import (
	"fmt"
	inter "sistema-clp/interface"
	"sistema-clp/structs"
)

func main() {
	inter.Menu()

	var prod structs.Produto
	var item structs.ItemVenda

	prod.NovoProduto(0, "teste", 10.5)

	item.NovoItemVenda(prod, 10.5, 2)

	fmt.Println(item.Produto.Nome)
}
