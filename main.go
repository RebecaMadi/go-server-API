package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)


func main() {
	//Cria o roteador
	router := mux.NewRouter()
	spa := spaHandler{staticPath: "src", indexPath: "index.html"}
	//router.PathPrefix("/").Handler(spa)
	
	HandlersHandle(router, spa)
	/*func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}*/
	
	//Dados do servidor
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//Encerra a execução se tiver erro
	log.Fatal(srv.ListenAndServe())
}