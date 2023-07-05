package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s \n ", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static")) // := means declare and define
	http.Handle("/", fileServer) 
	http.HandleFunc("/form", formHandler) // pull formhandler function when accessing /form 
	http.HandleFunc("/hello", helloHandler) // pull hellohandler function when accesing /hello

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{ // if err listenandserve :8080 //nil means null // if function in go is like this
		log.Fatal(err) // produce error when err is not equal nil 
	}
	
}