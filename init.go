package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile(serviceKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app : %v\n", err)
	}
	// [END intialize_app_service_account_golang]
	return app
}
