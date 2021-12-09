package inter

import (
	"fmt"
)

func Menu() {
	var opt int

	fmt.Println("Seja bem vindo!")
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Inserir clinte")
	fmt.Println("2 - Inserir produto")
	fmt.Println("3 - Registrar venda")
	fmt.Println("4 - Sair")

	fmt.Scanf("%d", &opt)

	switch opt {
	case 1:
	}
}
