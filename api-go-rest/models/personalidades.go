package models //Importando pacote models

type Personalidade struct { //Criando estrutura de dados
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade //Atribuindo variável
