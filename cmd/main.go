package main

import (
	"fmt"
	"os"

	"github.com/diegofilbal/gobank/cmd/bank"
)

func main() {

	banco := bank.Banco{}

	for {
		fmt.Println("\n======================================")
		fmt.Println("                GOBANK                ")
		fmt.Println("======================================")
		fmt.Println("Selecione uma opção:")
		fmt.Println("1. Cadastrar Conta Normal")
		fmt.Println("2. Cadastrar Conta Bônus")
		fmt.Println("3. Consultar Saldo")
		fmt.Println("4. Realizar Crédito")
		fmt.Println("5. Realizar Débito")
		fmt.Println("6. Realizar Transferência")
		fmt.Println("7. Render Juros")
		fmt.Println("8. Sair")
		fmt.Println("======================================")
		fmt.Print("Digite uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			numeroConta := bank.SolicitarNumeroConta()
			banco.CriarConta(numeroConta, bank.CONTA_NORMAL)
		case 2:
			numeroConta := bank.SolicitarNumeroConta()
			banco.CriarConta(numeroConta, bank.CONTA_BONUS)
		case 3:
			numeroConta := bank.SolicitarNumeroConta()
			banco.ConsultarSaldo(numeroConta)
		case 4:
			numeroConta := bank.SolicitarNumeroConta()
			valor := bank.SolicitarValor()
			banco.RealizarCredito(numeroConta, valor)
		case 5:
			numeroConta := bank.SolicitarNumeroConta()
			valor := bank.SolicitarValor()
			banco.RealizarDebito(numeroConta, valor)
		case 6:
			numeroContaOrigem := bank.SolicitarNumeroContaOrigem()
			numeroContaDestino := bank.SolicitarNumeroContaDestino()
			valor := bank.SolicitarValor()
			banco.RealizarTransferencia(numeroContaOrigem, numeroContaDestino, valor)
		case 7:
			taxa := bank.SolicitarTaxa()
			banco.RenderJuros(taxa)
		case 8:
			fmt.Println("Saindo da aplicação...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Por favor, selecione novamente.")
		}
	}
}
