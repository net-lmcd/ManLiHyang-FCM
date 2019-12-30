package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("/Users/user/firebase/manlihyang-55bb9-firebase-adminsdk-jicmz-02b377bcfc.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app : %v\n", err)
	}
	// [END intialize_app_service_account_golang]
	return app
}
