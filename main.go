package main

import (
	"fmt"
	"os"
)

func main() {
	//go build - arquivo main
	//go run - arquivo main
	intro()

	showMenu()

	option := getOption()
	switch option {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Fechando o programa...")
		os.Exit(0) // funcao para fechar um programa
	default:
		fmt.Println("Não conheço este comando.")

	}

}

func intro() {
	nome := "Jocel"
	fmt.Println("Olá, sr.", nome)
}
func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

// passando retorno da variavel depois da declaração da funcao
func getOption() int {
	var optionRead int
	//get int
	fmt.Scan(&optionRead)
	fmt.Println("O comando escolhido foi", optionRead)
	return optionRead
}
