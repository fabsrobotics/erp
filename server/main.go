package main

import (
	"fmt"
	"net/http"
	"os"
)

func MainRouter(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
}

func main(){
	fmt.Println("Server initialized on port "+os.Getenv("PORT"))
	// Main Router
	http.HandleFunc("/",MainRouter)
	// Listen And Serve
	err := http.ListenAndServe(":"+os.Getenv("PORT"),nil)
	if err != nil { panic(err) }
}