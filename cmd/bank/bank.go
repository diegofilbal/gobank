package bank

import (
	"fmt"
	"regexp"
)

const (
	CONTA_NORMAL   = "Normal"
	CONTA_BONUS    = "Bônus"
	CONTA_POUPANCA = "Poupança"
)

type Conta struct {
	numero    int
	saldo     float64
	pontuacao int
	tipo			string
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

func SolicitarNumeroContaOrigem() int {
	var numero int
	fmt.Print("Digite o número da conta de origem: ")
	fmt.Scanln(&numero)
	return numero
}

func SolicitarNumeroContaDestino() int {
	var numero int
	fmt.Print("Digite o número da conta de destino: ")
	fmt.Scanln(&numero)
	return numero
}

func SolicitarTaxa() float64 {
	var numero float64
	fmt.Print("Digite a taxa de juros (%) que deseja render nas contas poupança cadastradas: ")
	fmt.Scanln(&numero)
	return numero
}

func SolicitarValor() float64 {
	var valor float64
	fmt.Print("Digite o valor: ")
	fmt.Scanln(&valor)
	return valor
}

func solicitarSaldoInicial() float64 {
	var valor float64
	fmt.Print("Digite o saldo inicial: ")
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

func saldoSuficiente(conta Conta, valor float64) bool {
	return valor <= conta.saldo
}

func (c *Conta) imprime() {
	fmt.Println("\n--------------------------------------")
	fmt.Println("            Dados da Conta            ")
	fmt.Println("--------------------------------------")
	fmt.Printf("Número: %d\n", c.numero)
	fmt.Printf("Tipo: %s\n", c.tipo)
	fmt.Printf("Saldo: %.2f\n", c.saldo)
	if c.tipo == CONTA_BONUS {
		fmt.Printf("Pontuação: %d\n", c.pontuacao)
	}
}

func (b *Banco) CriarConta(numero int, tipoConta string) {
	if !numeroContaValido(numero) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}

	conta := b.buscaConta(numero)
	if conta == nil {
		saldoInicial := 0.0
		if tipoConta == CONTA_NORMAL {
			saldoInicial = solicitarSaldoInicial()
			if !valorValido(saldoInicial) {
				fmt.Println("Valor inválido. Certifique-se de que seja um número real positivo.")
				return
			}
		}
		novaConta := Conta{numero: numero, saldo: saldoInicial, tipo: tipoConta}
		if tipoConta == CONTA_BONUS {
			novaConta.pontuacao = 10
		}

		b.contas = append(b.contas, novaConta)
		fmt.Println("Conta criada com sucesso!")
		novaConta.imprime()
	} else {
		fmt.Printf("Já existe conta para número %d. Tente outro número.\n", numero)
	}
}

func (b *Banco) ConsultarSaldo(numero int) {
	if !numeroContaValido(numero) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}
	conta := b.buscaConta(numero)
	if conta != nil {
		fmt.Println("Conta encontrada.")
		conta.imprime()
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
		if conta.tipo == CONTA_BONUS {
			conta.pontuacao += int(valor / 100)
		}
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
		if saldoSuficiente(*conta, valor) {
			conta.saldo -= valor
			fmt.Printf("Débito de %.2f realizado com sucesso. Novo saldo: %.2f\n", valor, conta.saldo)
		} else {
			fmt.Println("Saldo insuficiente para realizar o débito.")
		}
	} else {
		fmt.Printf("Conta %d não encontrada\n", numero)
	}
}

func (b *Banco) RealizarTransferencia(numeroOrigem int, numeroDestino int, valor float64) {
	if !numeroContaValido(numeroOrigem) || !numeroContaValido(numeroDestino) {
		fmt.Println("Número de conta inválido. Certifique-se de que seja um número inteiro positivo.")
		return
	}
	if numeroOrigem == numeroDestino {
		fmt.Println("Conta de origem e destino não podem ser iguais.")
		return
	}
	contaOrigem := b.buscaConta(numeroOrigem)
	if contaOrigem == nil {
		fmt.Printf("Conta %d não encontrada\n", numeroOrigem)
		return
	}
	contaDestino := b.buscaConta(numeroDestino)
	if contaDestino == nil {
		fmt.Printf("Conta %d não encontrada\n", numeroDestino)
		return
	}
	if !valorValido(valor) {
		fmt.Println("Valor inválido. Certifique-se de que seja um número real positivo.")
		return
	}
	if saldoSuficiente(*contaOrigem, valor) {
		contaOrigem.saldo -= valor
		contaDestino.saldo += valor
		if contaDestino.tipo == CONTA_BONUS {
			contaDestino.pontuacao += int(valor / 200)
		}
		fmt.Printf("Transferência de %.2f realizada com sucesso.\n", valor)
		fmt.Printf("Novo saldo da conta %d: %.2f\n", numeroOrigem, contaOrigem.saldo)
		fmt.Printf("Novo saldo da conta %d: %.2f\n", numeroDestino, contaDestino.saldo)
	} else {
		fmt.Println("Saldo insuficiente na conta de origem para realizar a transferência.")
	}
}

func (c *Conta) renderJuros(taxaJuros float64) {
	if c.tipo == CONTA_POUPANCA {
		c.saldo *= (1 + taxaJuros/100)
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
	fmt.Println("Juros aplicados em todas as contas poupança cadastradas com sucesso.")
}
