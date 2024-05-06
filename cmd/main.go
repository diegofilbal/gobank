package main

import (
	"fmt"
	"os"

	"github.com/diegofilbal/gobank/cmd/bank"
)

func main() {
	for {
		fmt.Println("Selecione uma opção:")
		fmt.Println("1. Cadastrar Conta")
		fmt.Println("2. Consultar Saldo")
		fmt.Println("3. Realizar Crédito")
		fmt.Println("4. Realizar Débito")
		fmt.Println("5. Realizar Transferência")
		fmt.Println("6. Sair")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			banco := bank.Banco{}
			numeroConta := bank.SolicitarNumeroConta()
			banco.CriarConta(numeroConta)
		case 2:
			fmt.Println("Consultar Saldo")
		case 3:
			fmt.Println("Realizar Crédito")
		case 4:
			fmt.Println("Realizar Débito")
		case 5:
			fmt.Println("Realizar Transferência")
		case 6:
			fmt.Println("Sair")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Por favor, selecione novamente.")
		}
	}
}
