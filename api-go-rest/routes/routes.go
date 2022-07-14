package routes //Chamando pacote de rotas

import ( //Importando bibliotecas
	"api-go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() { //Chamando função de HandleRequest
	r := mux.NewRouter()                                                                         //Criando nova rota com gorilla mux
	r.HandleFunc("/", controllers.Home)                                                          //Chamando função de imprimir na tela do site
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")          //Imprimindo na tela todas as personalidades
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get") //Imprimindo na tela uma personalidade específica
	log.Fatal(http.ListenAndServe(":8000", r))                                                   //Passando qual porta devemos ouvir e servir em caso de log.fatal
}
