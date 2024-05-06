package bank

import (
	"fmt"
	"regexp"
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

func (b *Banco) buscaConta(numero int) *Conta {
	for _, conta := range b.contas {
		if conta.numero == numero {
			return &conta
		}
	}

	return nil
}

func numeroContaValido(numero int) bool {
	if numero <= 0 {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+$", fmt.Sprint(numero))
	return match
}

func (b *Banco) CriarConta(numero int) {
	if !numeroContaValido(numero) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}
	conta := b.buscaConta(numero)
	if conta == nil {
		saldoInicial := 0.0
		novaConta := Conta{numero: numero, saldo: saldoInicial}
		b.contas = append(b.contas, novaConta)
		fmt.Printf("Conta criada com sucesso: número %d, saldo inicial %.2f\n", numero, saldoInicial)
	} else {
		fmt.Printf("Já existe conta para número %d. Tente outro numero.\n", numero)
	}

}

func (b* Banco) ConsultarSaldo(numero int) {
	if !numeroContaValido(numero) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}
	conta := b.buscaConta(numero)
	if conta != nil {
		fmt.Printf("Conta %d encontrada. Saldo: %.2f\n", numero, conta.saldo)
	} else {
		fmt.Printf("Conta %d não encontrada\n", numero)
	}
}