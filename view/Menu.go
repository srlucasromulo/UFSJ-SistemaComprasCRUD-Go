package view

import "fmt"

func Menu() int {

	fmt.Println("|-----MENU INICIAL-----|")
	fmt.Println("1 - Menu cliente")
	fmt.Println("2 - Menu produto")
	fmt.Println("3 - Menu venda")
	fmt.Println("0 - Sair")
	fmt.Println("|----------------------|")

	var opt int
	_, err := fmt.Scanf("%d", &opt)
	if err != nil {
		fmt.Println("Algo de errado não está certo.")
		return 0
	}

	return opt
}
