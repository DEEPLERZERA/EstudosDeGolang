package main //Cria biblioteca

import (
	"Net/http" //Importa biblioteca Net/http
	"fmt"      //Importa biblioteca fmt
	"os"       //Importa biblioteca de os
)

func main() { //Criando função principal
	for { //Faz com que se repita infinitamente
		ExibeIntroducao() //Chamando função
		ExibeMenu()       //Chamando função

		comando := LeComando() //Atribuindo função a comando

		/*if comando == 1 {    //Criando condicionais
			fmt.Println("Monitorando.......") //Imprimindo na tela
		} else if comando == 2 {  //Mais condicionais
			fmt.Println("Exibindo logs.....")  //Imprime na tela
		}else if comando == 0 {  //Condicional
			fmt.Println("Saindo do programa....")  //Imprime na tela
		}else {  //Se não faça
			fmt.Println("Não reconheço esse comando.....")  //Imprima na tela
		}
		*/

		switch comando { //Criando switch-case
		case 1: //Caso seja 1 faça
			IniciarMonitoramento() //Chamando função
		case 2: //Caso seja 2 faça
			fmt.Println("Exibindo logs.....") //Imprime na tela
		case 0: //Caso seja 0 faça
			fmt.Println("Saindo do programa....") //Imprime na tela
			os.Exit(0)                            //Sai do programa sem erro
		default: //Se não for nenhuma das opções a cima faça isto
			fmt.Println("Não reconheço esse comando.....") //Imprima na tela
			os.Exit(-1)                                    //Sai do programa com erro
		}
	}

}

func ExibeIntroducao() { //Cria função de exibir introducao
	nome := "Daniel" //Criando variáveis
	versao := 1.1
	fmt.Println("Hello World!") //Imprimindo na tela
	fmt.Println("Olá Senhor ", nome)
	fmt.Println("Estamos na versão: ", versao)
}

func ExibeMenu() { //Cria funcao de exibir menu
	fmt.Println("1 - Iniciar ") //Imprimindo menu na tela
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

}

func LeComando() int { //Cria funcao de ler comando
	var comandoLido int    //Criando variavel de comando lido
	fmt.Scan(&comandoLido) //Recebendo input do usuário

	fmt.Println("O comando escolhido foi", comandoLido) //Imprimindo na tela

	return comandoLido //Dando retorno pois é uma função que exige retornar valores
}

func IniciarMonitoramento() {
	fmt.Println("Monitorando.......") //Imprimindo na tela
	site := "https://www.youtube.com" //Atribuindo site a variável
	resp, _ := http.Get(site)         //Pegando resposta http do site e atribuindo a variável

	if resp.StatusCode == 200 { //Verifica se acesso ao site teve exito
		fmt.Println("Site:", site, "foi carregado com sucesso!") //Imprime na tela que sim
	} else { //Se não faz
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode) //Imprime na tela que deu erro e imprime o statuscode do site
	}
}
