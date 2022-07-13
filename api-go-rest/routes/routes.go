package routes //Chamando pacote de rotas

import ( //Importando bibliotecas
	"api-go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() { //Chamando função de HandleRequest
	http.HandleFunc("/", controllers.Home)       //Chamando função de imprimir na tela do site
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)   //Imprimindo na tela todas as personalidades
	log.Fatal(http.ListenAndServe(":8000", nil)) //Passando qual porta devemos ouvir e servir em caso de log.fatal
}
