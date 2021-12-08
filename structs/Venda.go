package structs

import "time"

type Venda struct {
	numero  int
	data    time.Time
	cliente Cliente
	itens   []ItemVenda
}

func (v *Venda) NovaVenda(numero int, data time.Time, cliente Cliente) {
	v.numero = numero
	v.data = data
	v.cliente = cliente
}
