package contas

import (
	"go_oo/clientes"
)

type ContaPoupanca struct {
	Titular                        clientes.Titular
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeScar := valorSaque <= c.saldo && valorSaque > 0

	if podeScar {
		c.saldo -= valorSaque
		return "Saque sucess"
	} else {
		return "Saque insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "sucess", c.saldo
	} else {
		return "fail", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
