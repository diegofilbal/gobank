package bank

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	CONTA_NORMAL   = "Normal"
	CONTA_BONUS    = "Bônus"
	CONTA_POUPANCA = "Poupança"
)

type Conta struct {
	Numero    int     `json:"numero"`
	Saldo     float64 `json:"saldo"`
	Pontuacao int     `json:"pontuacao,omitempty"`
	Tipo      string  `json:"tipo"`
}

type Banco struct {
	contas []Conta
}

func numeroContaValido(numero int) bool {
	if numero <= 0 {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+$", fmt.Sprint(numero))
	return match
}

func valorValido(valor float64) bool {
	if !(valor > 0) {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+(\\.[0-9]+)?$", fmt.Sprint(valor))
	return match
}

func saldoSuficiente(conta Conta, valor float64) bool {
	return valor <= conta.Saldo
}

func (b *Banco) buscaConta(numero int) *Conta {
	for i := range b.contas {
		if b.contas[i].Numero == numero {
			return &b.contas[i]
		}
	}

	return nil
}

func (b *Banco) CriarConta(numero int, tipoConta string, saldoInicial float64) error {
	if !numeroContaValido(numero) {
		return errors.New("Número de conta inválido")
	}

	conta := b.buscaConta(numero)
	if conta != nil {
		return errors.New("Conta já existe")
	}

	saldo := 0.0

	if tipoConta == CONTA_NORMAL {
		if !valorValido(saldoInicial) {
			return errors.New("Saldo inicial inválido")
		}
		saldo = saldoInicial
	}

	novaConta := Conta{
		Numero: numero,
		Tipo:   tipoConta,
		Saldo:  saldo,
	}

	if tipoConta == CONTA_BONUS {
		novaConta.Pontuacao = 10
	}

	b.contas = append(b.contas, novaConta)
	return nil
}

func (b *Banco) ConsultarConta(numero int) *Conta {
	if !numeroContaValido(numero) {
		return nil
	}
	return b.buscaConta(numero)
}

func (b *Banco) ConsultarSaldo(numero int) *Conta {
	if !numeroContaValido(numero) {
		return nil
	}
	return b.buscaConta(numero)
}

func (b *Banco) RealizarCredito(numero int, valor float64) error {
	if !numeroContaValido(numero) {
		return errors.New("Número de conta inválido")
	}
	conta := b.buscaConta(numero)
	if conta == nil {
		return errors.New("Conta não encontrada")
	}
	if !valorValido(valor) {
		return errors.New("Valor inválido")
	}

	conta.Saldo += valor
	if conta.Tipo == CONTA_BONUS {
		conta.Pontuacao += int(valor / 100)
	}

	return nil
}

func (b *Banco) RealizarDebito(numero int, valor float64) error {
	if !numeroContaValido(numero) {
		return errors.New("Número de conta inválido")
	}
	conta := b.buscaConta(numero)
	if conta == nil {
		return errors.New("Conta não encontrada")
	}
	if !valorValido(valor) {
		return errors.New("Valor inválido")
	}
	if !saldoSuficiente(*conta, valor) {
		return errors.New("Saldo insuficiente")
	}

	conta.Saldo -= valor
	return nil
}

func (b *Banco) RealizarTransferencia(numeroOrigem int, numeroDestino int, valor float64) error {
	if !numeroContaValido(numeroOrigem) || !numeroContaValido(numeroDestino) {
		return errors.New("Número de conta inválido")
	}
	if numeroOrigem == numeroDestino {
		return errors.New("Conta de origem e destino não podem ser iguais")
	}
	contaOrigem := b.buscaConta(numeroOrigem)
	if contaOrigem == nil {
		return errors.New("Conta de origem não encontrada")
	}
	contaDestino := b.buscaConta(numeroDestino)
	if contaDestino == nil {
		return errors.New("Conta de destino não encontrada")
	}
	if !valorValido(valor) {
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
