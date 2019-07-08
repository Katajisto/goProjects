package main

import (
	"net/http"
	"goProjects/main-page"
	"goProjects/portfolio"
	"log"
)


func main() {
	http.HandleFunc("/", mp.Handle)
	http.HandleFunc("/portfolio/",portfolio.ServeHome)
	http.HandleFunc("/portfolio/text/",portfolio.ServePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
