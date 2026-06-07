package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Cambio struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

var ErrArgumentosInsuficientes = errors.New(
	"uso: ./conversor <valor> <moeda>\nexemplo: ./conversor 100 USD",
)

type ErrValorInvalido struct {
	Valor string
}

func (e *ErrValorInvalido) Error() string {
	return fmt.Sprintf("'%s' não é um número válido", e.Valor)
}

type ErrMoedaNaoEncontrada struct {
	Moeda string
}

func (e *ErrMoedaNaoEncontrada) Error() string {
	return fmt.Sprintf("moeda '%s' não encontrada", e.Moeda)
}

func parsearArgumentos() (string, string, error) {
	if len(os.Args) < 3 {
		return "", "", ErrArgumentosInsuficientes
	}
	return os.Args[1], os.Args[2], nil
}

func parsearValor(rawValor string) (float64, error) {
	valor, err := strconv.ParseFloat(rawValor, 64)

	if err != nil {
		return 0, &ErrValorInvalido{Valor: rawValor}
	}
	return valor, nil
}

func buscarTaxa(cambio Cambio, moeda string) (float64, error) {
	taxa, ok := cambio.Rates[moeda]
	if !ok {
		return 0, &ErrMoedaNaoEncontrada{Moeda: moeda}
	}
	return taxa, nil
}

func main() {
	dados := `{
		"base": "BRL",
		"date": "2025-04-14",
		"rates": {
			"USD": 0.151,
			"EUR": 0.137,
			"JPY": 16.29,
			"GBP": 0.13,
			"CHF": 0.1402,
			"AUD": 0.2712,
			"CAD": 0.2374,
			"CNY": 1.251,
			"HKD": 1.326,
			"NZD": 0.2922,
			"SEK": 1.655,
			"NOK": 1.806,
			"DKK": 1.122,
			"SGD": 0.2249,
			"KRW": 242.97,
			"ZAR": 3.239,
			"MXN": 3.454,
			"INR": 14.71,
			"ILS": 0.63,
			"THB": 5.74,
			"IDR": 2875.0,
			"MYR": 0.754,
			"PHP": 9.74,
			"PLN": 0.644,
			"CZK": 3.77,
			"HUF": 61.59,
			"TRY": 6.49,
			"BGN": 0.293,
			"RON": 0.746
		}
	}
	`

	cambio := Cambio{}

	if err := json.Unmarshal([]byte(dados), &cambio); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "erro ao carregar taxas:", err)
		os.Exit(1)
	}

	rawValor, moeda, err := parsearArgumentos()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	valor, err := parsearValor(rawValor)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	taxa, err := buscarTaxa(cambio, moeda)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%.2f\n", valor*taxa)

}
