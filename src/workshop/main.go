package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	//"bytes"
)

import "github.com/cdipaolo/sentiment"

func main() {
	// create our json file
	path := "response.json"
	DeleteFile(path)
	CreateFile(path)

	// create our sentiment analysis model
	model, err := sentiment.Restore()

	// make sure model was restored correctly
	checkError(err)

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
			fmt.Print(analysis.Score)

			// jsonify our data
			data, err := json.Marshal(analysis)
			checkError(err)

			// convert json []byte type to a string
			stringifiedJson := string(data)
			// write our data to json file
			WriteFile(path, stringifiedJson)
			// separately, return json output to user
			json.NewEncoder(w).Encode(analysis)
		}
	})

	// start up the server
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func CreateFile(path string) {
	// check to see if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

// auxiliary function which handles error checking for us
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func WriteFile(path string, data string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	// defer will always be run as the last line of this function
	defer file.Close()

	// write our text to file
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

func DeleteFile(path string) {
	// delete file
	os.Remove(path)
}
