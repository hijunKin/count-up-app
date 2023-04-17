package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
  	fmt.Println("start")

  	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatalf("can't read the file, err = %v", err)
		return 
	}

	htmlStr = string(data)

	http.HandleFunc("/", handler) 

	http.HandleFunc("/countup", countUpHandler)

	http.HandleFunc("/getindexhtml", getIndexHTML)

	http.ListenAndServe(":8081", nil)
}

var htmlStr string
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, count)
	fmt.Fprintln(w, count)
} 

func countUpHandler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
} 

func getIndexHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(htmlStr))
} 