package structs

import "time"

type Cliente struct {
	Pessoa
	rg   string
	bday time.Time
}

func (c *Cliente) NovoCliente(nome string, endereco string, rg string, bday time.Time) {
	c.SetNome(nome)
	c.SetEndereco(endereco)
	c.SetRg(rg)
	c.SetBday(bday)
}

func (c *Cliente) GetNome() string {
	return c.nome
}

func (c *Cliente) SetNome(nome string) {
	c.nome = nome
}

func (c *Cliente) GetEndereco() string {
	return c.endereco
}

func (c *Cliente) SetEndereco(endero string) {
	c.endereco = endero
}

func (c *Cliente) GetRg() string {
	return c.rg
}

func (c *Cliente) SetRg(rg string) {
	c.rg = rg
}

func (c *Cliente) GetBday() time.Time {
	return c.bday
}

func (c *Cliente) SetBday(bday time.Time) {
	c.bday = bday
}
