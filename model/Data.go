package model

type Data struct {
	dia int
	mes int
	ano int
}

func (b *Data) NovaData(dia int, mes int, ano int) {
	b.SetDia(dia)
	b.SetMes(mes)
	b.SetAno(ano)
}

func (b *Data) SetDia(dia int) {
	b.dia = dia
}

func (b *Data) SetMes(mes int) {
	b.mes = mes
}

func (b *Data) SetAno(ano int) {
	b.ano = ano
}

func (b *Data) GetDia() int {
	return b.dia
}

func (b *Data) GetMes() int {
	return b.mes
}

func (b *Data) GetAno() int {
	return b.ano
}
