package controllers //Chamando pacote de controllers

import ( //Importando bibliotecas
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) { //Criando função home que consegue escrever e fazer request
	fmt.Fprint(w, "Home Page") //Imprime na página
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) { //Criando função que escreve e faz request no http
	json.NewEncoder(w).Encode(models.Personalidades) //Chamando models.Personalidade e aplicando encode
}
