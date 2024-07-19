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
		fmt.Println("3. Cadastrar Conta Poupança")
		fmt.Println("4. Consultar Conta")
		fmt.Println("5. Consultar Saldo")
		fmt.Println("6. Realizar Crédito")
		fmt.Println("7. Realizar Débito")
		fmt.Println("8. Realizar Transferência")
		fmt.Println("9. Render Juros")
		fmt.Println("10. Sair")
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
			banco.CriarConta(numeroConta, bank.CONTA_POUPANCA)
		case 4:
			numeroConta := bank.SolicitarNumeroConta()
			banco.ConsultarConta(numeroConta)
		case 5:
			numeroConta := bank.SolicitarNumeroConta()
			banco.ConsultarSaldo(numeroConta)
		case 6:
			numeroConta := bank.SolicitarNumeroConta()
			valor := bank.SolicitarValor()
			banco.RealizarCredito(numeroConta, valor)
		case 7:
			numeroConta := bank.SolicitarNumeroConta()
			valor := bank.SolicitarValor()
			banco.RealizarDebito(numeroConta, valor)
		case 8:
			numeroContaOrigem := bank.SolicitarNumeroContaOrigem()
			numeroContaDestino := bank.SolicitarNumeroContaDestino()
			valor := bank.SolicitarValor()
			banco.RealizarTransferencia(numeroContaOrigem, numeroContaDestino, valor)
		case 9:
			taxa := bank.SolicitarTaxa()
			banco.RenderJuros(taxa)
		case 10:
			fmt.Println("Saindo da aplicação...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Por favor, selecione novamente.")
		}
	}
}
