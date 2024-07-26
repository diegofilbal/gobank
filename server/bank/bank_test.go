package bank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCadastrarConta(t *testing.T) {
	tests := []struct {
		numero   int
		tipo     string
		saldo    float64
		expected Conta
	}{
		{
			numero: 1,
			tipo:   CONTA_NORMAL,
			saldo:  100.0,
			expected: Conta{
				Numero:    1,
				Tipo:      CONTA_NORMAL,
				Saldo:     100.0,
				Pontuacao: 0,
			},
		},
		{
			numero: 2,
			tipo:   CONTA_BONUS,
			saldo:  10.0,
			expected: Conta{
				Numero:    2,
				Tipo:      CONTA_BONUS,
				Saldo:     0.0,
				Pontuacao: 10,
			},
		},
		{
			numero: 3,
			tipo:   CONTA_POUPANCA,
			saldo:  500.0,
			expected: Conta{
				Numero:    3,
				Tipo:      CONTA_POUPANCA,
				Saldo:     0.0,
				Pontuacao: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.tipo, func(t *testing.T) {
			b := Banco{}
			err := b.CriarConta(tt.numero, tt.tipo, tt.saldo)
			require.NoError(t, err)
			conta := b.BuscaConta(tt.numero)
			require.NotNil(t, conta)
			assert.Equal(t, tt.expected, *conta)
		})
	}
}

func TestConsultarConta(t *testing.T) {
	b := Banco{}
	err := b.CriarConta(1, CONTA_NORMAL, 100.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.CriarConta(2, CONTA_BONUS, 50.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.CriarConta(3, CONTA_POUPANCA, 200.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	tests := []struct {
		numero   int
		expected *Conta
	}{
		{
			numero: 1,
			expected: &Conta{
				Numero: 1,
				Tipo:   CONTA_NORMAL,
				Saldo:  100.0,
			},
		},
		{
			numero: 2,
			expected: &Conta{
				Numero:    2,
				Tipo:      CONTA_BONUS,
				Saldo:     0.0,
				Pontuacao: 10,
			},
		},
		{
			numero: 3,
			expected: &Conta{
				Numero:    3,
				Tipo:      CONTA_POUPANCA,
				Saldo:     0.0,
				Pontuacao: 0,
			},
		},
		{
			numero:   999,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run("ConsultarConta", func(t *testing.T) {
			conta := b.ConsultarConta(tt.numero)
			if tt.expected == nil {
				assert.Nil(t, conta)
			} else {
				require.NotNil(t, conta)
				assert.Equal(t, *tt.expected, *conta)
			}
		})
	}
}

func TestRealizarCredito(t *testing.T) {
	b := Banco{}
	err := b.CriarConta(1, CONTA_NORMAL, 100.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.CriarConta(2, CONTA_BONUS, 0.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	tests := []struct {
		numero   int
		valor    float64
		expected *Conta
		errMsg   string
	}{
		{
			numero:   1,
			valor:    50.0,
			expected: &Conta{Numero: 1, Tipo: CONTA_NORMAL, Saldo: 150.0},
			errMsg:   "",
		},
		{
			numero:   1,
			valor:    -50.0,
			expected: nil,
			errMsg:   "Valor inválido",
		},
		{
			numero:   2,
			valor:    100.0,
			expected: &Conta{Numero: 2, Tipo: CONTA_BONUS, Saldo: 100.0, Pontuacao: 11},
			errMsg:   "",
		},
		{
			numero:   999,
			valor:    50.0,
			expected: nil,
			errMsg:   "Conta não encontrada",
		},
	}

	for _, tt := range tests {
		t.Run("RealizarCredito", func(t *testing.T) {
			err := b.RealizarCredito(tt.numero, tt.valor)
			if tt.errMsg != "" {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				conta := b.BuscaConta(tt.numero)
				require.NotNil(t, conta)
				assert.Equal(t, *tt.expected, *conta)
			}
		})
	}
}

func TestRealizarDebito(t *testing.T) {
	b := Banco{}
	err := b.CriarConta(1, CONTA_NORMAL, 100.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.CriarConta(2, CONTA_BONUS, 50.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	tests := []struct {
		numero   int
		valor    float64
		expected *Conta
		errMsg   string
	}{
		{
			numero:   1,
			valor:    50.0,
			expected: &Conta{Numero: 1, Tipo: CONTA_NORMAL, Saldo: 50.0},
			errMsg:   "",
		},
		{
			numero:   1,
			valor:    -50.0,
			expected: nil,
			errMsg:   "Valor inválido",
		},
		{
			numero:   1,
			valor:    100.0,
			expected: nil,
			errMsg:   "Saldo insuficiente",
		},
		{
			numero:   999,
			valor:    50.0,
			expected: nil,
			errMsg:   "Conta não encontrada",
		},
	}

	for _, tt := range tests {
		t.Run("RealizarDebito", func(t *testing.T) {
			err := b.RealizarDebito(tt.numero, tt.valor)
			if tt.errMsg != "" {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				conta := b.BuscaConta(tt.numero)
				require.NotNil(t, conta)
				assert.Equal(t, *tt.expected, *conta)
			}
		})
	}
}

func TestRealizarTransferencia(t *testing.T) {
	b := Banco{}
	err := b.CriarConta(1, CONTA_NORMAL, 100.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.CriarConta(2, CONTA_BONUS, 0.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	tests := []struct {
		numeroOrigem    int
		numeroDestino   int
		valor           float64
		expectedOrigem  *Conta
		expectedDestino *Conta
		errMsg          string
	}{
		{
			numeroOrigem:    1,
			numeroDestino:   2,
			valor:           50.0,
			expectedOrigem:  &Conta{Numero: 1, Tipo: CONTA_NORMAL, Saldo: 50.0},
			expectedDestino: &Conta{Numero: 2, Tipo: CONTA_BONUS, Saldo: 50.0, Pontuacao: 10},
			errMsg:          "",
		},
		{
			numeroOrigem:    1,
			numeroDestino:   2,
			valor:           -50.0,
			expectedOrigem:  nil,
			expectedDestino: nil,
			errMsg:          "Valor inválido",
		},
		{
			numeroOrigem:    1,
			numeroDestino:   2,
			valor:           100.0,
			expectedOrigem:  nil,
			expectedDestino: nil,
			errMsg:          "Saldo insuficiente",
		},
		{
			numeroOrigem:    1,
			numeroDestino:   999,
			valor:           50.0,
			expectedOrigem:  nil,
			expectedDestino: nil,
			errMsg:          "Conta de destino não encontrada",
		},
		{
			numeroOrigem:    1,
			numeroDestino:   1,
			valor:           50.0,
			expectedOrigem:  nil,
			expectedDestino: nil,
			errMsg:          "Conta de origem e destino não podem ser iguais",
		},
	}

	for _, tt := range tests {
		t.Run("RealizarTransferencia", func(t *testing.T) {
			err := b.RealizarTransferencia(tt.numeroOrigem, tt.numeroDestino, tt.valor)
			if tt.errMsg != "" {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				origem := b.BuscaConta(tt.numeroOrigem)
				destino := b.BuscaConta(tt.numeroDestino)
				require.NotNil(t, origem)
				require.NotNil(t, destino)
				assert.Equal(t, *tt.expectedOrigem, *origem)
				assert.Equal(t, *tt.expectedDestino, *destino)
			}
		})
	}
}

func TestRenderJuros(t *testing.T) {
	b := Banco{}
	err := b.CriarConta(1, CONTA_POUPANCA, 0.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.RealizarCredito(1, 1000.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.CriarConta(2, CONTA_POUPANCA, 2000.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	err = b.RealizarCredito(2, 2000.0)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	tests := []struct {
		taxa      float64
		expected1 float64
		expected2 float64
	}{
		{
			taxa:      10.0,
			expected1: 1100.0,
			expected2: 2200.0,
		},
		{
			taxa:      5.0,
			expected1: 1155.0,
			expected2: 2310.0,
		},
	}

	for _, tt := range tests {
		t.Run("RenderJuros", func(t *testing.T) {
			b.RenderJuros(tt.taxa)
			conta1 := b.BuscaConta(1)
			conta2 := b.BuscaConta(2)
			require.NotNil(t, conta1)
			require.NotNil(t, conta2)
			assert.Equal(t, tt.expected1, conta1.Saldo)
			assert.Equal(t, tt.expected2, conta2.Saldo)
		})
	}
}
