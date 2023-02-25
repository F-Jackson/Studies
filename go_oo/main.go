package main

import (
	"fmt"

	"go_oo/clientes"
	contas "go_oo/contas"
)

func PagarBoleto(conta verificarConta, valoBoleto float64) {
	conta.Sacar(valoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// Titular := "Guilherme"
	// NumAgencia := 589
	// NumConta := 123456
	// Saldo := 125.50
	clienteBruno := clientes.Titular{
		Nome:      "Guilherme",
		CPF:       "123.111.123.12",
		Profissao: "Dev"}

	contaDoBruno := contas.ContaCorrent{}
	contaDoBruno.Titular = clienteBruno
	contaDoBruno.NumAgencia = 589
	contaDoBruno.NumConta = 123456
	contaDoBruno.Depositar(325.50)

	// contaDaBruna := contas.ContaCorrent{"Bruna", 222, 111222, 100}

	// var contaDaCris *contas.ContaCorrent = new(contas.ContaCorrent)
	// contaDaCris.Titular = "Cris"

	fmt.Println(contaDoBruno)

	fmt.Println(contaDoBruno.Sacar(300))
	fmt.Println(contaDoBruno.Sacar(-300))
}
