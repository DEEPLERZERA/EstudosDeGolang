package main  //Chamando pacote principal

import (  //Importando bibliotecas

	"fmt"
	"bufio"    
	"os"
	"strings"
	"io"
	"math/rand"



)

var palavra_escolhida string   //Criando variáveis

func main() {  //Chamando função principal
//selecionando a palavra secreta
	palavra_escolhida := escolhe_palavra()  //Atribue palavra escolhida para a variável

//que comece o jogo
	jogo_da_forca(palavra_escolhida)  //Depois executa jogo da forca

	
	

}


func abre_arquivo() []string {   //Criando função de abrir arquivo
	var palavras []string //Chamando var palavras

	arquivo, err := os.Open("palavras.txt") //Abrindo arquivo de palavras e atribuindo a arquivo

	if err != nil { //Se houver erro faça
		fmt.Println("Ocorreu um erro:", err) //Imprime na tela mensagem de erro
	}

	leitor := bufio.NewReader(arquivo) //Lê linha a linha do arquivo de texto e atribue a variável leitor

	for { //Criando laço for de repetição
		linha, err := leitor.ReadString('\n') //Fazendo leitura individual de cada linha
		linha = strings.TrimSpace(linha)      //Tirando os \n
		palavras = append(palavras, linha)          //Atribuindo strings a slice

		if err == io.EOF { //Se for fim de arquivo
			break //Saia do loop for
		}
	}

	arquivo.Close() //Feche o arquivo aberto a cima(Questão de boa prática)

	return palavras //Retornando palavras
}


func escolhe_palavra() string{  //Criando função de escolher palavras
	conteudo := abre_arquivo()  //Atribue retorno da função abre_arquivo() a conteudo
	randomIndex := rand.Intn(len(conteudo))  //Randomizando o conteudo
	pick := conteudo[randomIndex]  //Fazendo um pick aleatorio 
	return pick   //Retornando pick
}


func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}




