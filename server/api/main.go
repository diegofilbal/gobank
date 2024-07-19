package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/diegofilbal/gobank/bank"
	"github.com/gin-gonic/gin"
)

var banco = bank.Banco{}

func main() {
	r := gin.Default()

	r.POST("banco/conta", cadastrarConta)
	r.GET("banco/conta/:numero", consultarConta)
	r.GET("banco/conta/:numero/saldo", consultarSaldo)
	r.PUT("banco/conta/:numero/credito", realizarCredito)
	r.PUT("banco/conta/:numero/debito", realizarDebito)
	r.PUT("/banco/conta/transferencia", realizarTransferencia)
	r.PUT("/banco/conta/rendimento", renderJuros)

	err := r.Run(":8001")
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

func cadastrarConta(ctx *gin.Context) {
	var req struct {
		Numero int     `json:"numero"`
		Tipo   string  `json:"tipo"`
		Saldo  float64 `json:"saldoInicial"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	banco.CriarConta(req.Numero, req.Tipo, req.Saldo)
	ctx.JSON(http.StatusOK, gin.H{"message": "Conta criada com sucesso"})
}

func consultarConta(ctx *gin.Context) {
	numero, err := strconv.Atoi(ctx.Param("numero"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Número de conta inválido"})
		return
	}

	conta := banco.ConsultarConta(numero)
	if conta == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Conta não encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"saldo": conta.Saldo, "tipo": conta.Tipo, "pontuacao": conta.Pontuacao})
}

func consultarSaldo(ctx *gin.Context) {
	numero, err := strconv.Atoi(ctx.Param("numero"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Número de conta inválido"})
		return
	}

	conta := banco.ConsultarConta(numero)
	if conta == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Conta não encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"conta": numero, "saldo": conta.Saldo})
}

func realizarCredito(ctx *gin.Context) {
	var req struct {
		Valor float64 `json:"valor"`
	}

	numero, err := strconv.Atoi(ctx.Param("numero"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Número de conta inválido"})
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	if err := banco.RealizarCredito(numero, req.Valor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Crédito realizado com sucesso"})
}

func realizarDebito(ctx *gin.Context) {
	var req struct {
		Valor float64 `json:"valor"`
	}

	numero, err := strconv.Atoi(ctx.Param("numero"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Número de conta inválido"})
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	if err := banco.RealizarDebito(numero, req.Valor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Débito realizado com sucesso"})
}

func realizarTransferencia(ctx *gin.Context) {
	var req struct {
		NumeroOrigem  int     `json:"from"`
		NumeroDestino int     `json:"to"`
		Valor         float64 `json:"amount"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	if err := banco.RealizarTransferencia(req.NumeroOrigem, req.NumeroDestino, req.Valor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transferência realizada com sucesso"})
}

func renderJuros(ctx *gin.Context) {
	var req struct {
		Taxa float64 `json:"taxa"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	banco.RenderJuros(req.Taxa)
	ctx.JSON(http.StatusOK, gin.H{"message": "Juros aplicados com sucesso"})
}
