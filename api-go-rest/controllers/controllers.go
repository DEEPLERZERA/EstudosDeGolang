package controllers //Chamando pacote de controllers

import ( //Importando bibliotecas
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) { //Criando função home que consegue escrever e fazer request
	fmt.Fprint(w, "Home Page") //Imprime na página
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) { //Criando função que escreve e faz request no http
	var p []models.Personalidade //Criando slice p que recebe personalidades
	database.DB.Find(&p)         //O banco de dados faz uma varredura
	json.NewEncoder(w).Encode(p) //Chamando models.Personalidade e aplicando encode
}

//CRIANDO FUNÇÃO QUE RETORNA A PERSONALIDADE DE ACORDO COM ID PASSADO PELO USUÁRIO NA URL
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade   //Variável personalidade recebe os modelos
	database.DB.Find(&personalidade, id)     //Database acha a personalidade de acordo com id
	json.NewEncoder(w).Encode(personalidade) //Imprime na tela os dados

}
