package bank

import (
	"fmt"
)

type Conta struct {
	numero int
	saldo  float64
}

type Banco struct {
	contas []Conta
}

func SolicitarNumeroConta() int {
	var numero int
	fmt.Println("Digite o número da conta:")
	fmt.Scanln(&numero)
	return numero
}

func (b *Banco) CriarConta(numero int) {
	saldoInicial := 0.0
	novaConta := Conta{numero: numero, saldo: saldoInicial}
	b.contas = append(b.contas, novaConta)
	fmt.Printf("Conta criada com sucesso: numero %d, saldo inicial %.2f\n", numero, saldoInicial)
}
