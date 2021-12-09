package structs

import "time"

type Venda struct {
	numero  int
	data    time.Time
	cliente Cliente
	itens   []ItemVenda
}

func (v *Venda) NovaVenda(numero int, data time.Time, cliente Cliente) {
	v.SetNumero(numero)
	v.SetData(data)
	v.SetCliente(cliente)
}

func (v *Venda) GetNumero() int {
	return v.numero
}

func (v *Venda) SetNumero(numero int) {
	v.numero = numero
}

func (v *Venda) GetData() time.Time {
	return v.data
}

func (v *Venda) SetData(data time.Time) {
	v.data = data
}

func (v *Venda) GetCliente() Cliente {
	return v.cliente
}

func (v *Venda) SetCliente(cliente Cliente) {
	v.cliente = cliente
}

func (v *Venda) Total() float32 {
	var soma float32 = 0
	for i := range v.itens {
		soma += v.itens[i].Total()
	}
	return soma
}
