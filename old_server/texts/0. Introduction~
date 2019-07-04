package main

import(
	"net/http"
	"log"
	"time"
	"fmt"
)

type pageData struct {
	Title       string
	Text        string
	TimeToServe string
}

type mainData struct {
	Texts []text
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	data := mainData{
		Texts: textSlice,
	}
	mainTemplatePtr.Execute(w, data)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	keys, ok := r.URL.Query()["text"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Parameter missing")
		return
	}
	key := keys[0]
	var queryText text
	var emptyText text
	for i := 0; i < len(textSlice); i++ {
		if textSlice[i].Title == key {
			queryText = textSlice[i]
			break
		}
	}
	if queryText == emptyText {
		log.Println("Text doesnt exist")
		return
	}
	data := pageData{
		Title: queryText.Title,
		Text: queryText.Text,
		TimeToServe: fmt.Sprintf("%v",time.Since(start)),
	}
	templatePtr.Execute(w, data)
}
