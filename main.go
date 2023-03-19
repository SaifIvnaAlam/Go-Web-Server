package main

import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "PareForm() : err : %v", err)
		return
	}
	fmt.Fprintf( w, "Post Request Sucessful")
	name := r.FormValue("name")
	food := r.FormValue( "food")
	fmt.Fprintf(w,"Your Name is %s\n", name)
	fmt.Fprintf(w, "Your Favourite Food is%s\n", food)
}

func helloHandler(w http.ResponseWriter,r *http.Request ){
if r.URL.Path !="/hello"{ http.Error(w, "404 Not Found",http.StatusNotFound )
return
}
if r.Method !="Get"{
	http.Error(w, "Method Not Found",http.StatusNotFound)
	return
}
fmt.Fprintf(w, "Hello Amgios!")
}
func main()  {
	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err:= http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
	
}