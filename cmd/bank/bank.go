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
	fmt.Print("Digite o número da conta: ")
	fmt.Scanln(&numero)
	return numero
}

func SolicitarValor() float64 {
	var valor float64
	fmt.Print("Digite o valor: ")
	fmt.Scanln(&valor)
	return valor
}

func (b *Banco) buscaConta(numero int) *Conta {
	for i := range b.contas {
		if b.contas[i].numero == numero {
			return &b.contas[i]
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

func valorValido(valor float64) bool {
	if valor <= 0 {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+(\\.[0-9]+)?$", fmt.Sprint(valor))
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

func (b *Banco) ConsultarSaldo(numero int) {
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

func (b *Banco) RealizarCredito(numero int, valor float64) {
	if !numeroContaValido(numero) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}
	conta := b.buscaConta(numero)
	if conta != nil {
		if !valorValido(valor) {
			fmt.Println("Valor inválido. Certifique-se de que seja um número real positivo.")
			return
		}
		conta.saldo += valor
		fmt.Printf("Crédito de %.2f realizado com sucesso. Novo saldo: %.2f\n", valor, conta.saldo)
	} else {
		fmt.Printf("Conta %d não encontrada\n", numero)
	}
}

func (b *Banco) RealizarDebito(numero int, valor float64) {
	if !numeroContaValido(numero) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}
	conta := b.buscaConta(numero)
	if conta != nil {
		if !valorValido(valor) {
			fmt.Println("Valor inválido. Certifique-se de que seja um número real positivo.")
			return
		}
		conta.saldo -= valor
		fmt.Printf("Débito de %.2f realizado com sucesso. Novo saldo: %.2f\n", valor, conta.saldo)
	} else {
		fmt.Printf("Conta %d não encontrada\n", numero)
	}
}
