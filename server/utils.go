package utils

import (
	"fmt"
	"regexp"
)

func NumeroContaValido(numero int) bool {
	if numero <= 0 {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+$", fmt.Sprint(numero))
	return match
}

func ValorValido(valor float64) bool {
	if valor <= 0 {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+(\\.[0-9]+)?$", fmt.Sprint(valor))
	return match
}

func LimiteSaldoNegativo(valor float64, saldo float64) bool {
	limite := saldo - valor

	if limite < -2000 {
		return true
	} else {
		return false
	}
}
