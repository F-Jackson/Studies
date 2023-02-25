package contas

import (
	"go_oo/clientes"
)

type ContaCorrent struct {
	Titular    clientes.Titular
	NumAgencia int
	NumConta   int
	saldo      float64
}

func (c *ContaCorrent) Sacar(valorSaque float64) string {
	podeScar := valorSaque <= c.saldo && valorSaque > 0

	if podeScar {
		c.saldo -= valorSaque
		return "Saque sucess"
	} else {
		return "Saque insuficiente"
	}
}

func (c *ContaCorrent) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "sucess", c.saldo
	} else {
		return "fail", c.saldo
	}
}

func (c *ContaCorrent) Transferir(valor float64, contaDestino *ContaCorrent) bool {
	if valor <= c.saldo && valor > 0 {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrent) ObterSaldo() float64 {
	return c.saldo
}
