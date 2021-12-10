package main

import (
	"fmt"
	inter "sistema-clp/interface"
	structs "sistema-clp/structs"
	//"time"
)

func main() {

	for {
		opt := inter.Menu()

		switch opt {

		case 1, 2:
			crud := inter.MenuCRUD(opt)
			fmt.Println(crud)

		}

		if opt == 0 {
			break
		}
	}

}

func imprimeTotal(modulo structs.Totalizavel) {
	fmt.Println(modulo.Total())
}
