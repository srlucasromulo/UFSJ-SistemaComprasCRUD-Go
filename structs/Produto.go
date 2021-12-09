package structs

type Produto struct {
	codigo int
	nome   string
	valor  float32
}

func (p *Produto) NovoProduto(codigo int, nome string, valor float32) {
	p.SetCodigo(codigo)
	p.SetNome(nome)
	p.SetValor(valor)
}

func (p *Produto) GetCodigo() int {
	return p.codigo
}

func (p *Produto) SetCodigo(codigo int) {
	p.codigo = codigo
}

func (p *Produto) GetNome() string {
	return p.nome
}

func (p *Produto) SetNome(nome string) {
	p.nome = nome
}

func (p *Produto) GetValor() float32 {
	return p.valor
}

func (p *Produto) SetValor(valor float32) {
	p.valor = valor
}
