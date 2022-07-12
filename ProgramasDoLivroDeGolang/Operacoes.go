package main  //Chamando pacote principal

import "fmt"  //Importando biblioteca fmt

type Operacao interface {  //Criando tipo operacao interface
	Calcular() int  //Chamando função Calcular int
}

type Soma struct {  //Criando tipo Soma que é uma struct
	operando1, operando2 int  //Declarando valores int
}

func (s Soma) Calcular() int {  //Criando função de somar
	return s.operando1 + s.operando2  //Retornando resultado
}

func (s Soma) String() string {  //Criando função que imprime
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)  //Imprimindo e concatenando valores
}


type Subtracao struct {  // Criando tipo subtracao que é uma struct
	operando1, operando2 int //Declarando variáveis 
}

func (s Subtracao) Calcular() int {  //Criando função de calcular subtracao
	return s.operando1 - s.operando2  ///Retornando resultado
}

func (s Subtracao) String() string {  //Imprimindo na tela
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2) //Imprimindo e concatenando valores
}

func main() {  //Chamando função principal
	operacoes := make([]Operacao, 4)  //Criando slice e atribuindo a operações
	operacoes[0] = Soma{10, 20}  //Fazendo operações e atribuindo a slice operacoes
	operacoes[1] = Subtracao{30, 15}
	operacoes[2] = Subtracao{10, 50}
	operacoes[3] = Soma{5, 2}

	acumulador := 0  //Declarando acumulador

	for _, op := range operacoes {  //Varrendo operacoes com op
		valor := op.Calcular()  //Atribuindo a variável valor op.Calcular()
		fmt.Printf("%v = %d\n", op, valor)  //Imprimindo na tela
		acumulador += valor  //Aumentando valor de acumulador
	}
	fmt.Println("Valor acumulado =", acumulador)  //Imprimindo na tela
}


