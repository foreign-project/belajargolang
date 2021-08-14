package main

import (
	"log"
	"net/http"
) 

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("Stating web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to home"))
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world, saya sedang belajar Golang web"))
}

func marioHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Mario from Nintendo"))
}