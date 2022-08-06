package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
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
	fmt.Printf("# of CPU: %v\n", runtime.NumCPU())
	reqCount += 1
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	fmt.Printf("%v requested\n", reqCount)
	time.Sleep(5 * time.Second)
	fmt.Println("Processed...")
	fmt.Fprintln(w, "Hello!")
}

var reqCount = 0

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
