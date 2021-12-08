package structs

type ItemVenda struct {
	produto    Produto
	valor      float32
	quantidade int
}

func (i *ItemVenda) NovoItemVenda(p Produto, valor float32, quantidade int) {
	i.produto = p
	i.valor = valor
	i.quantidade = quantidade
}
