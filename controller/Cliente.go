package controller

import (
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
)

func NovoCliente(clientes []structs.Cliente) []structs.Cliente {
	nome, endereco, rg, _, _, _ := inter.MenuNovoCliente()
	var c structs.Cliente
	c.NovoCliente(nome, endereco, rg)
	clientes = append(clientes, c)
	return clientes
}

func AtualizaCliente(rg string, clientes []structs.Cliente) {
	for i := range clientes {
		if clientes[i].GetRg() == rg {
			nome, endereco := inter.MenuAtualizaCliente()
			clientes[i].SetEndereco(endereco)
			clientes[i].SetNome(nome)
		}
	}
}

func ApagarCliente(rg string, clientes []structs.Cliente) []structs.Cliente {
	var index int
	for i := range clientes {
		if clientes[i].GetRg() == rg {
			index = i
			break
		}
	}

	clientes = append(clientes[:index], clientes[index+1:]...)

	return clientes
}
