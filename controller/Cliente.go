package controller

import (
	"fmt"
	"sistema-clp/model"
	"sistema-clp/view"
)

func NovoCliente() model.Cliente {

	nome, endereco, rg, _, _, _ := view.MenuNovoCliente()
	var c model.Cliente
	c.NovoCliente(nome, endereco, rg)
	return c
}

func AdicionarNovoCliente(c model.Cliente, clientes []model.Cliente) []model.Cliente {

	clientes = append(clientes, c)
	return clientes
}

func AtualizaCliente(clientes []model.Cliente) []model.Cliente {

	view.MenuListaCliente(clientes)

	var rg string
	fmt.Println("Digite o RG do cliente: ")
	fmt.Scanf("%s", &rg)

	for i := range clientes {
		if clientes[i].GetRg() == rg {
			nome, endereco, rg, _, _, _ := view.MenuNovoCliente()
			var c model.Cliente
			c.NovoCliente(nome, endereco, rg)
			clientes[i] = c
		}
	}
	return clientes
}

func ApagarCliente(clientes []model.Cliente) []model.Cliente {

	view.MenuListaCliente(clientes)

	var rg string
	fmt.Println("Digite o RG do cliente: ")
	fmt.Scanf("%s", &rg)

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
