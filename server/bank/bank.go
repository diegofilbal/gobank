package bank

import (
	"errors"
	"fmt"

	utils "github.com/diegofilbal/gobank"
)

const (
	CONTA_NORMAL   = "Normal"
	CONTA_BONUS    = "Bonus"
	CONTA_POUPANCA = "Poupança"
)

type Conta struct {
	Numero    int
	Saldo     float64
	Pontuacao int
	Tipo      string
}

type Banco struct {
	contas []Conta
}

func (b *Banco) BuscaConta(numero int) *Conta {
	for i := range b.contas {
		if b.contas[i].Numero == numero {
			return &b.contas[i]
		}
	}
	return nil
}

func saldoSuficiente(conta Conta, valor float64) bool {
	return valor <= conta.Saldo
}

func (b *Banco) CriarConta(numero int, tipoConta string, saldoInicial float64) error {
	if !utils.NumeroContaValido(numero) {
		return errors.New("Número de conta inválido")
	}

	conta := b.BuscaConta(numero)
	if conta != nil {
		return errors.New("Conta já existe")
	}

	novaConta := Conta{
		Numero: numero,
		Tipo:   tipoConta,
		Saldo:  saldoInicial,
	}

	if tipoConta == CONTA_BONUS {
		novaConta.Pontuacao = 10
	}

	if tipoConta != CONTA_NORMAL {
		novaConta.Saldo = 0
	}

	b.contas = append(b.contas, novaConta)
	return nil
}

func (b *Banco) ConsultarConta(numero int) *Conta {
	if !utils.NumeroContaValido(numero) {
		return nil
	}
	return b.BuscaConta(numero)
}

func (b *Banco) ConsultarSaldo(numero int) *Conta {
	if !utils.NumeroContaValido(numero) {
		return nil
	}
	return b.BuscaConta(numero)
}

func (b *Banco) RealizarCredito(numero int, valor float64) error {
	if !utils.NumeroContaValido(numero) {
		return errors.New("Número de conta inválido")
	}
	conta := b.BuscaConta(numero)
	if conta == nil {
		return errors.New("Conta não encontrada")
	}
	if !utils.ValorValido(valor) {
		return errors.New("Valor inválido")
	}

	// conta.Saldo += valor
	if conta.Tipo == CONTA_BONUS {
		conta.Pontuacao += int(valor / 100)
	}

	return nil
}

func (b *Banco) RealizarDebito(numero int, valor float64) error {
	if !utils.NumeroContaValido(numero) {
		return errors.New("Número de conta inválido")
	}
	conta := b.BuscaConta(numero)
	if conta == nil {
		return errors.New("Conta não encontrada")
	}
	if !utils.ValorValido(valor) {
		return errors.New("Valor inválido")
	}
	if !saldoSuficiente(*conta, valor) {
		return errors.New("Saldo insuficiente")
	}

	conta.Saldo -= valor
	return nil
}

func (b *Banco) RealizarTransferencia(numeroOrigem int, numeroDestino int, valor float64) error {
	if !utils.NumeroContaValido(numeroOrigem) || !utils.NumeroContaValido(numeroDestino) {
		return errors.New("Número de conta inválido")
	}
	if numeroOrigem == numeroDestino {
		return errors.New("Conta de origem e destino não podem ser iguais")
	}
	contaOrigem := b.BuscaConta(numeroOrigem)
	if contaOrigem == nil {
		return errors.New("Conta de origem não encontrada")
	}
	contaDestino := b.BuscaConta(numeroDestino)
	if contaDestino == nil {
		return errors.New("Conta de destino não encontrada")
	}

	if !utils.ValorValido(valor) {
		return errors.New("Valor inválido")
	}

	if !saldoSuficiente(*contaOrigem, valor) {
		return errors.New("Saldo insuficiente")
	}

	contaOrigem.Saldo -= valor
	contaDestino.Saldo += valor
	if contaDestino.Tipo == CONTA_BONUS {
		contaDestino.Pontuacao += int(valor / 150)
	}

	return nil
}

func (c *Conta) renderJuros(taxaJuros float64) {
	if c.Tipo == CONTA_POUPANCA {
		c.Saldo *= (1 + taxaJuros/100)
	}
}

func (b *Banco) RenderJuros(taxaJuros float64) {
	if len(b.contas) == 0 {
		fmt.Println("Operação cancelada, o banco ainda não possui nenhuma conta cadastrada.")
		return
	}
	for i := range b.contas {
		b.contas[i].renderJuros(taxaJuros)
	}
}
