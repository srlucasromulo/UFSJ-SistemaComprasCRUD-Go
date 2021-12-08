package structs

type Produto struct {
	codigo int
	nome   string
	valor  float32
}

func (p *Produto) NovoProduto(codigo int, nome string, valor float32) {
	p.codigo = codigo
	p.nome = nome
	p.valor = valor
}
