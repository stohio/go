# Golang Workshop
# Motivation
## Why Go? 
Go is the hip new language. It has all the ease of use of a language such as Python, but still can give you nitty gritty control that C++ gives. It is especially optimized for use as a server side programming language - that is why Software Lab is built in Go! 
## fileops.go
fileops.go is full of common file operations that are each wrapped into a function, for ease of use. 
## sentiment.go
sentiment.go loads up a prebuilt sentiment analysis model (naive bayes), and performs analysis on any string. 
## main.go
main.go puts together all the pieces, to create a RESTful API which performs sentiment analysis on any string sent by GET request and returns the analysis via JSON. 
```
go run main.go

Paste this into your Chrome Browser! 
http://localhost:8080/i%20love%20you!%20I%20always%20love%20this%20hackathon
```
