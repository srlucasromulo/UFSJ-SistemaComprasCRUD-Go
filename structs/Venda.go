package structs

type Venda struct {
	produto    Produto
	valor      float32
	quantidade int
	itens      []ItemVenda
}

//func (v Venda) Total() {
//
//}
