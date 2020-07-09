package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoThales := contas.ContaCorrente{Titular: "Thales", Saldo: 300}
	contaDoArthur := contas.ContaCorrente{Titular: "Arthur", Saldo: 100}

	fmt.Println(contaDoThales)
	fmt.Println(contaDoArthur)

	status := contaDoArthur.Transferir(-200, &contaDoThales)

	fmt.Println(status)
	fmt.Println(contaDoThales)
	fmt.Println(contaDoArthur)
}
