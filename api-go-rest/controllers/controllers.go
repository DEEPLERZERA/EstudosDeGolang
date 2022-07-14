package controllers //Chamando pacote de controllers

import ( //Importando bibliotecas
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) { //Criando função home que consegue escrever e fazer request
	fmt.Fprint(w, "Home Page") //Imprime na página
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) { //Criando função que escreve e faz request no http
	json.NewEncoder(w).Encode(models.Personalidades) //Chamando models.Personalidade e aplicando encode
}

//CRIANDO FUNÇÃO QUE RETORNA A PERSONALIDADE DE ACORDO COM ID PASSADO PELO USUÁRIO NA URL
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades { //Percorrendo todos modelos de Personalidades
		if strconv.Itoa(personalidade.Id) == id { //Caso seja igual ao id passado pelo usuário faça
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
