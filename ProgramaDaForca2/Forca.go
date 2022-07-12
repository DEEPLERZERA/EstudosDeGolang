package main  //Importando pacote principal

import (  //Importando bibliotecas
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strings"
	"io"

)

func contem(s string, caracteres []string) bool {  //Criando função pra verificar se contem
	for _, ca := range s {  //Percorrendo s com ca
		for _, ca2 := range caracteres {  //Percorrendo caracteres que foi recebido como parametro por ca2
			if string(ca) == ca2 {  //Fazendo uma comparação
				return true  //Retornando verdadeiro
			}
		}
	}
	return false  //Retornando falso
}

func entrada(strings []string, separator string) string {  //Função que recebe entrada
	if len(strings) == 0 {  //Verifica se tamanho das strings é igual a 0
		return ""  //Retorna string vazia
	}
	s := ""   //Atribue vazio a s
	ultimoIdx := len(strings) - 1 //Verifica qual é o último index
	for _, v := range strings[0:ultimoIdx] {  //Percorre o slice strings com v
		s += v + separator  //Vai concatenando
	}
	return s + strings[ultimoIdx]  //Retorne a string
}

func main() {  //Função principal
	t := time.Now()  //Atribue tempo de agora
	rand.Seed(t.UnixNano())  //Gera um seed aleatória a cada nano segundo
	palavra_escolhida := escolhePalavra()  //Atribue palavra escolhida para a variável
	nGuesses := len(palavra_escolhida)  //Pegando tamanho da palavra e atribuindo a nGuesses
	encontrado := []string{} //Criando slice de encontrado

	for i := 0; i < len(palavra_escolhida); i++ {  //Criando laço de repetição
		encontrado = append(encontrado, "_") //Vai adicionando _ _ _ _ a tela do usuário
	}
	for nGuesses > 0 {  //Para tamanho da palavra maior que zero
		fmt.Println("Você tem", nGuesses, "Chances sobrando!")  //Mostra número de chances sobrando
		letra, err := pegaLetra(encontrado)  //Chamando função de pegar letra e atribuindo a letra
		if err != nil {  //Se for nulo faça
			fmt.Println("Erro de leitura no console!")  //Imprime erro na tela
			return
		}
		if !contem(palavra_escolhida, []string{letra}) {  //Se palavra_escolhida não estiver contida em string letras faça
			nGuesses--  //Diminue nGuesses
		}
		if atualizandoEncontrado(encontrado, palavra_escolhida, letra) {  //Se encontrar palavra escolhida
			fmt.Println("Você ganhou ganhou! a palavra era:", palavra_escolhida)  //Imprime na tela
			return  //Da retorno
		}
	}
	fmt.Println("Você perdeu! a palavra era:", palavra_escolhida)  //Imprime na tela
}

func pegaLetra(encontrou []string) (string, error) {  //Criando função de pegar letra
	ALFABETO := "aáàãbcçdeéêfghiíîjklmnoôõópqrstuúvwxyz"  //Atribuindo alfabeto

	for true {  //Laço de retição infinito
		letra, err := prompt("Escolha uma letra:", entrada(encontrou, " "))  //Recebe input do usuário e atribue a letra
		if err != nil {  //Se erro for diferente de nulo faça
			return "", err  //Retorne erro
		}
		if len(letra) == 1 && contem(ALFABETO, []string{letra}) {  //Condicional que verifica se está tudo certo com a letra entregue pelo usuário
			return letra, nil  //Retorna letra
		}
		fmt.Println("Input invalido: Tente com uma única letra minúsucula!")  //Se não imprime input invalido

	}
	return "", nil  //Retorna nulo
}


func atualizandoEncontrado(encontrado []string, palavra string, letra string) bool {  //Criando função de atualizar encontrado
	completo := true   //Completo passa a ser true
	for i, r := range palavra {  //Percorrendo palavra com r faça
		if letra == string(r) {  //Se letra for igual a r faça
			encontrado[i] = letra  //Atribue letra a encontrado
		}
		if encontrado[i] == "_"  {  //Se encontrado for igual a _ faça
			completo = false  //Completo ser false
		}

	}
	return completo //Retorne completo
}


func abreArquivo() []string {   //Criando função de abrir arquivo
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


func escolhePalavra() string{  //Criando função de escolher palavras
	conteudo := abreArquivo()  //Atribue retorno da função abre_arquivo() a conteudo
	indexAleatorio := rand.Intn(len(conteudo))  //Randomizando o conteudo
	escolha2 := conteudo[indexAleatorio]  //Fazendo um escolha2 aleatorio 
	return escolha2   //Retornando escolha2
}

func prompt(vals ...interface{}) (string, error) {  //Criando função de prompt
	if len(vals) != 0 {  //Se tamanho de vals for diferente de zero faça
		fmt.Println(vals...)  //imprima vals
	}
	leitura := bufio.NewScanner(os.Stdin)  //leitura recebe bufio.NewScanner
	leitura.Scan()  //Vai lendo input
	err := leitura.Err()   //Se tiver erro atribue a err
	if err != nil {  //Se tiver erro
		return "", err  //Imprime erro
	}
	
	return leitura.Text(), nil  //Retorna o texto scanneado e nil se tiver
}