package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	//"github.com/gorilla/websocket"

	"github.com/marcinpiecha/covid19/handlers"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	//http.HandleFunc("/favicon.ico", http.NotFound)
	r.HandleFunc("/favicon.ico", http.NotFound)
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/{key}", handlers.Individual)
	//http.ListenAndServe(":8080", r)
	fmt.Println("Applicatoin runnig on: http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
