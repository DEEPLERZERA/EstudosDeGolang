package main


import "fmt"

type ContaCorrente struct {
	titular string 
	numeroAgencia int 
	numeroConta int 
	saldo float64
}


func main() {
	contaDoGuilherme := ContaCorrente {"Guilherme", 589, 123456, 125.5}
	contaDaBruna := ContaCorrente {"Bruna", 222, 111222, 200}
	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
}
