package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteThales := clientes.Titular{"Thales", "123.111.123.12", "Programador"}
	contaDoThales := contas.ContaCorrente{Titular: clienteThales, NumeroAgencia: 123, NumeroConta: 123456}
	contaDoThales.Depositar(100)

	PagarBoleto(&contaDoThales, 60)

	fmt.Println(contaDoThales)

	contaDoArthur := contas.ContaPoupanca{}
	contaDoArthur.Depositar(320)

	PagarBoleto(&contaDoArthur, 120)

	fmt.Println(contaDoArthur)
}
