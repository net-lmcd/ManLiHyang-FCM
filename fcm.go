package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func sendToken(app *firebase.App) bool {
	/*
		START send_to_token_golang
		Obtain a messaging.Client from App
	*/
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("ERROR GETTING MESSAGING CLIENT: %v\n", err)
		return false

	}
	registrationToken := "hU6DCNDkGGBn40Dixd0hFy2m6bvm-7g0VCIq2n09_QXdpWxCy6Z"
	message := &messaging.Message{
		Data:  creatMessage("TITLE", "MESSAGE"),
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	fmt.Println("Successfully sent message : ", response)
	return true
}

func creatMessage(title string, body string) map[string]string {
	var message = make(map[string]string)
	message["title"] = title
	message["body"] = body
	return message
}
