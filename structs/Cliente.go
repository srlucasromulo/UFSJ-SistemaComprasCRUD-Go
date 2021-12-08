package structs

import "time"

type Cliente struct {
	Pessoa
	rg   string
	bday time.Time
}

func (c *Cliente) NovoCliente(nome string, endereco string, rg string, bday time.Time) {
	c.nome = nome
	c.endereco = endereco
	c.rg = rg
	c.bday = bday
}
