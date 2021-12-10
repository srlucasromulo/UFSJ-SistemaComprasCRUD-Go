package model

type Bday struct {
	dia int
	mes int
	ano int
}

func (b *Bday) NovoBday(dia int, mes int, ano int) {
	b.SetDia(dia)
	b.SetMes(mes)
	b.SetAno(ano)
}

func (b *Bday) SetDia(dia int) {
	b.dia = dia
}

func (b *Bday) SetMes(mes int) {
	b.mes = mes
}

func (b *Bday) SetAno(ano int) {
	b.ano = ano
}

func (b *Bday) GetDia() int {
	return b.dia
}

func (b *Bday) GetMes() int {
	return b.mes
}

func (b *Bday) GetAno() int {
	return b.ano
}
