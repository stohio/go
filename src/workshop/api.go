package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

import "github.com/cdipaolo/sentiment"

func main() {
	fmt.Fprintf(w, "Hello, %q", query))
	
	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// remove the "/" character in the beginning of query before doing sentiment analysis
		query := html.EscapeString(r.URL.Path) 
		analysis := model.SentimentAnalysis("Your mother is an awful lady", sentiment.English) // 0
		fmt.Print(analysis);
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
