package main  //Cria biblioteca

import "fmt" //Importa fmt



func main() {  //Criando função principal
	nome   := "Daniel" //Criando variáveis
	versao := 1.1
	fmt.Println("Hello World!")  //Imprimindo na tela
	fmt.Println("Olá Senhor ", nome)
	fmt.Println("Estamos na versão: ", versao)
	
	fmt.Println("1 - Iniciar ")  //Imprimindo menu na tela
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int  //Declarando variável

	fmt.Scan(&comando)  //Recebendo input do usuário
	
	fmt.Println("O comando escolhido foi", comando)  //Imprimindo na tela
	
}