package model

type ItemVenda struct {
	produto    Produto
	valor      float32
	quantidade int
}

func (item *ItemVenda) NovoItemVenda(p Produto, valor float32, quantidade int) {
	item.SetProduto(p)
	item.SetValor(valor)
	item.SetQuantidade(quantidade)
}

func (item *ItemVenda) GetProduto() Produto {
	return item.produto
}

func (item *ItemVenda) SetProduto(produto Produto) {
	item.produto = produto
}

func (item *ItemVenda) GetValor() float32 {
	return item.valor
}

func (item *ItemVenda) SetValor(valor float32) {
	item.valor = valor
}

func (item *ItemVenda) GetQuantidade() int {
	return item.quantidade
}

func (item *ItemVenda) SetQuantidade(quantidade int) {
	item.quantidade = quantidade
}

func (item *ItemVenda) Total() float32 {
	return item.valor * float32(item.quantidade)
}
