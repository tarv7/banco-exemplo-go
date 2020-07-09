package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteThales := clientes.Titular{"Thales", "123.111.123.12", "Programador"}
	contaDoThales := contas.ContaCorrente{Titular: clienteThales, NumeroAgencia: 123, NumeroConta: 123456}
	contaDoThales.Depositar(100)

	fmt.Println(contaDoThales.ObterSaldo())
}
