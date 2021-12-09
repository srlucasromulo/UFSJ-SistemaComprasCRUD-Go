package structs

type ItemVenda struct {
	Produto    Produto
	valor      float32
	quantidade int
}

func (i *ItemVenda) NovoItemVenda(p Produto, valor float32, quantidade int) {
	i.Produto = p
	i.valor = valor
	i.quantidade = quantidade
}

func (i *ItemVenda) Total() float32 {
	return i.valor * float32(i.quantidade)
}
