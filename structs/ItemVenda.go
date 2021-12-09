package structs

type ItemVenda struct {
	Produto    Produto
	Valor      float32
	Quantidade int
}

func (i *ItemVenda) NovoItemVenda(p Produto, valor float32, quantidade int) {
	i.Produto = p
	i.Valor = valor
	i.Quantidade = quantidade
}

func (i *ItemVenda) Total() float32 {
	return i.Valor * float32(i.Quantidade)
}
