package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getMessage() string
}

type EnglishBot struct {
	Message string
}
type SpaniceBot struct {
	Message string
}

func main() {

	englishMessage := EnglishBot{
		Message: "Hello there",
	}
	spanicehMessage := EnglishBot{
		Message: "Hola",
	}
	printMsg(englishMessage)
	printMsg(spanicehMessage)

	//http req example
	makeHttpReq()

	//read file data
	readFileData()
}
func printMsg(msgBot bot) {
	fmt.Println(msgBot.getMessage())
}

func (msg EnglishBot) getMessage() string {
	return msg.Message
}
func (msg SpaniceBot) getMessage() string {
	return msg.Message
}

//make http request

func makeHttpReq() {
	resp, err := http.Get("http://dev-server.autonom8.com/kinara-test/#/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//first way
	// byte := make([]byte, 99999)
	// resp.Body.Read(byte)
	// fmt.Println(string(byte))

	//second way

	io.Copy(os.Stdout, resp.Body)
}
