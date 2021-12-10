package model

type Totalizavel interface {
	Total() float32
}

func ImprimeTotal(modulo Totalizavel) float32 {
	return modulo.Total()
}
