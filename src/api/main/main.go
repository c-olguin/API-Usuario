package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"../controllers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/registrar", controllers.CrearUsuario).Methods("POST")
	r.HandleFunc("/login", controllers.AutenticarUsuario).Methods("POST")

	r.HandleFunc("/usuarios", controllers.ObtenerUsuarios).Methods("GET")
	r.HandleFunc("/usuarios/{id}", controllers.ObtenerUsuario).Methods("GET")
	r.HandleFunc("/usuarios/{id}", controllers.ModificarUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{id}", controllers.EliminarUsuario).Methods("DELETE")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	log.Fatal(http.ListenAndServe(":7031", corsOpts.Handler(r)))
	http.Handle("/", r)
}

// func JwtVerify(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var header = r.Header.Get("x-access-token") //Grab the token from the header

// 		header = strings.TrimSpace(header)

// 		if header == "" {
// 			//Token is missing, returns with error code 403 Unauthorized
// 			w.WriteHeader(http.StatusForbidden)
// 			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
// 			return
// 		}
// 		tk := &models.Token{}

// 		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("secret"), nil
// 		})

// 		if err != nil {
// 			w.WriteHeader(http.StatusForbidden)
// 			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "user", tk)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
