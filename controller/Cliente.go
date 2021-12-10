package controller

import (
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
)

func NovoCliente(clientes []structs.Cliente) {

	nome, endereco, rg, dia, mes, ano := inter.MenuNovoCliente()
	var c structs.Cliente
	c.NovoCliente(nome, endereco, rg)
	// adicionar em clientes
}
