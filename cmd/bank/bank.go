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

func (b *Banco) existeConta(numero int) bool {
	for _, conta := range b.contas {
		if conta.numero == numero {
			return true
		}
	}

	return false
}

func (b *Banco) CriarConta(numero int) {
	if !b.existeConta(numero) {
		saldoInicial := 0.0
		novaConta := Conta{numero: numero, saldo: saldoInicial}
		b.contas = append(b.contas, novaConta)
		fmt.Printf("Conta criada com sucesso: número %d, saldo inicial %.2f\n", numero, saldoInicial)
	} else {
		fmt.Printf("Já existe conta para número %d. Tente outro numero.\n", numero)
	}

}
