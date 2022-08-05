package main

import (
	"fmt"
	"log"
	"net/http"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err)
		return
	}
	name := r.FormValue("name")
	zipcode := r.FormValue("zipcode")
	fmt.Printf("name: %v,", name)
	fmt.Printf("zipcode: %v\n", zipcode)
	fmt.Fprintln(w, "Successfully signed up!")
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/greeting", greetingHandler)
	fmt.Println("Server listening at the port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}

}
