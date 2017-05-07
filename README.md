# Intro to Golang
## Install
- https://golang.org/dl/
- Add the following to top of ~/.bash_profile: 
```
export GOPATH=$HOME/go
export GOBIN=$HOME/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
```
## Why us? 
Our names are Steven Gates (Junior, Applied Math & CS) and Cameron Sinko (Freshman, Computer Engineering). In January, we went to a hackathon called BrickHacks in Rochester, New York. There, we created Software Lab because we were planning to host a hackathon in our hometown in Akron, Ohio. Turns out, other hackathons had the same problem we were having so we got funded by a group called EXL Center at The University of Akron to travel around and help hackathons. 
## Why Go? 
Go is the hip new language. It has all the ease of use of a language such as Python, but still can give you nitty gritty control that C++ gives. It is especially optimized for use as a server side programming language - that is why Software Lab is built in Go! 

## Useful Commands
- go install
- go run <filename>
- go get <libname>


## What's Possible
- Structs
- Pointers
- Defer
- Libraries (even from github!)
- Dynamic type setting
- returning multiple things
- Test suites
- File manipulations


NLP api server, simple
Make a JSON file, put stuff in it, restful api, return them a JSON

Copy paste this into your Chrome browser to test out the sentiment analysis API! 
```
http://localhost:8081/?q=your+mother+is+an+awful+lady
```
