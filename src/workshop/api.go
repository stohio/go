package main

import (
	"fmt"
	"io"
	"html"
	"log"
	"net/http"
	"strings"
	"encoding/json"
	"os"
	//"bytes"

)

import "github.com/cdipaolo/sentiment"

func main() {
	// first time setup stuff
	
	// create our json file
	path := "response.json"
	deleteFile(path)
	createFile(path)

	// create our sentiment analysis model
	model, err := sentiment.Restore()

	// make sure model was restored correctly
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	// default router which handles our GET requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := html.EscapeString(r.URL.Path) 

		// clean and trim the query
		t := strings.Replace(query, "/", "", -1)

		// ignore favicon reqeusts (only happens with browser - because this is an MVP demo)
		if strings.TrimRight(t, "\n") != "favicon.ico" {
			// analyze our query
			analysis := model.SentimentAnalysis(t, sentiment.English) 
			// print the overall score for our convenience
			fmt.Print(analysis.Score);
			// return all data to client in form of a json
			data, err := json.Marshal(analysis)
			fmt.Print(err)
			s := string(data)
			
			writeFile(path, s)

			json.NewEncoder(w).Encode(analysis)
		}
	})

	// start up the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createFile(path string) {
	// check to see if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func writeFile(path string, data string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// write some text to file
	_, err = file.WriteString(data)
	checkError(err)

	// save changes
	err = file.Sync()
	checkError(err)
}

func readFile(path string) {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// read file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	fmt.Println(string(text))
	checkError(err)
}

func deleteFile(path string) {
	// delete file
	os.Remove(path)
}
