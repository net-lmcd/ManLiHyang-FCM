package main

import (
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
)

func sendMessage(request *Request) bool {
	/*
		START send_to_token_golang
		Obtain a messaging.Client from App
	*/
	message := &messaging.Message{
		Data:  converMessageToMap(&request.Message),
		Token: request.From,
	}

	log.Print("[SEND MESSAGE] : ", request)
	response, err := fcmClient.Send(fcmCtx, message)
	if err != nil {
		log.Fatalln("[SEND MESSAGE ERROR] : ", err.Error())
		return false
	}
	fmt.Println("Successfully sent message : ", response)
	return true
}

func converMessageToMap(message *Message) map[string]string {
	var result = make(map[string]string)
	result["title"] = message.Title
	result["body"] = message.Body
	return result
}

func creatMessage(title string, body string) map[string]string {
	var result = make(map[string]string)
	result["title"] = title
	result["body"] = body
	return result
}
