package structs

type Produto struct {
	codigo int
	Nome   string
	valor  float32
}

func (p *Produto) NovoProduto(codigo int, nome string, valor float32) {
	p.codigo = codigo
	p.Nome = nome
	p.valor = valor
}
