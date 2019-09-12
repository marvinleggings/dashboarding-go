package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type DashboardLayout struct{
	X int `json: "x"`
	Y int `json: "y"`
	Width int `json: "w"`
	height int `json: "h"`
	Content string `json: "content"`
}


func main() {

	log.Println("Starting file server")
	fs := http.FileServer(http.Dir("src/app/static"))

	// http.Handle("/", fs)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		switch r.Method {
		case "GET":
			http.ServeFile(w,r,"src/app/static")
		case "POST":
			b, _ := ioutil.ReadAll(r.Body)
			log.Printf("%v", b)
		default:
			fmt.Fprint(w, "Method not defined")
		}
	})

	mux.HandleFunc("/blah", func(w http.ResponseWriter, r *http.Request) {
		log.Println("fewfewf "+r.Method)
		fmt.Fprintln(w, fs)
	})

	log.Println("Listening on port 80")
	http.ListenAndServe(":80", mux)
}