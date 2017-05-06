# Golang Workshop File Details
## fileops.go
fileops.go is full of common file operations that are each wrapped into a function, for ease of use. 
## sentiment.go
sentiment.go loads up a prebuilt sentiment analysis model (naive bayes), and performs analysis on any string. 
## main.go
main.go puts together all the pieces, to create a RESTful API which performs sentiment analysis on any string sent by GET request and returns the analysis via JSON. 
```
go run main.go

// Paste this into your Chrome Browser! 
http://localhost:8081/i%20love%20you!%20I%20always%20love%20this%20hackathon
```
