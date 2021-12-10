package view

import "fmt"

func Menu() int {

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Menu cliente")
	fmt.Println("2 - Menu produto")
	fmt.Println("3 - Registrar venda")
	fmt.Println("0 - Sair")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		return 0
	}

	return opt
}
