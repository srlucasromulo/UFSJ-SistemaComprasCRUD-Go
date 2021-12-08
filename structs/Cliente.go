package structs

import "time"

type Cliente struct {
	Pessoa
	rg string
	bday time.Time
}
