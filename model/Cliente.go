package model

type Cliente struct {
	Pessoa
	rg   string
	bday Bday
}

//func (c *Cliente) NovoCliente(nome string, endereco string, rg string, bday time.Time) {
func (c *Cliente) NovoCliente(nome string, endereco string, rg string, b Bday) {
	c.SetNome(nome)
	c.SetEndereco(endereco)
	c.SetRg(rg)
	c.SetBday(b)
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

func (c *Cliente) GetBday() Bday {
	return c.bday
}

func (c *Cliente) SetBday(bday Bday) {
	c.bday = bday
}
