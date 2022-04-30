package main

import (
	"Project_2/client"
	"Project_2/models"
	jsoniter "github.com/json-iterator/go"
	"os"
)

func main() {
	content, err := os.ReadFile("GoLang_Test.txt")
	if err != nil {
		panic(err)
	}
	var textInputRequest models.TextInput
	textInputRequest.Text = string(content)
	output := client.RpcService(textInputRequest)
	outputStr, _ := jsoniter.Marshal(output)
	println(string(outputStr))
}
