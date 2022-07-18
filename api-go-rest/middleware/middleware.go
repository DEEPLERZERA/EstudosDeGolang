package middleware //Importando pacote middleware

//Importando bibliotecas
import "net/http"

//CRIANDO FUNÇÃO DE MIDDLEWARE
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json") //Aplicando json
		next.ServeHTTP(w, r)
	})

}
