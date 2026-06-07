package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Cambio struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
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

	err := json.Unmarshal([]byte(dados), &cambio)

	if err != nil {
		panic(err)
	}

	rawValor := os.Args[1]
	base := os.Args[2]

	valor, err := strconv.ParseFloat(rawValor, 64)

	if err != nil {
		panic(err)
	}

	taxa, ok := cambio.Rates[base]

	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "erro: moeda %q não encontrada\n", base)
		os.Exit(1)
	}

	fmt.Printf("%.2f\n", valor*taxa)

}
