package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	return strings.ToUpper(os.Args[1]), strings.ToUpper(os.Args[2]), nil
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
	dados, err := os.ReadFile("taxas.json")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	cambio := Cambio{}

	if err := json.Unmarshal(dados, &cambio); err != nil {
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
