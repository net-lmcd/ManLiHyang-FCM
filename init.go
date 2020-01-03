package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App
var fcmCtx = context.Background()
var fcmClient *messaging.Client

//init function
func init() {
	//[INIT START]
	log.Println("[INIT START]")
	getFirebaseInstance()
	log.Printf("[FIREBASE.APP INIT COMPLETED]")
	getFcmClient()
	log.Printf("[FCM CLIENT INIT COMPLETED]")
	initDatabase()
}

func getFirebaseInstance() *firebase.App {
	if firebaseApp == nil {
		firebaseApp = initializeAppWithServiceAccount()
	}
	return firebaseApp
}

func getFcmClient() *messaging.Client {
	if fcmClient == nil {
		fcmClient = initClient(getFirebaseInstance())
	}
	return fcmClient
}

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile(serviceKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app : %v\n", err)
	}
	// [END intialize_app_service_account_golang]
	return app
}

func initClient(app *firebase.App) *messaging.Client {
	client, err := app.Messaging(fcmCtx)
	if err != nil {
		log.Fatalf("ERROR GETTING MESSAGING CLIENT: %v\n", err)
		os.Exit(1)
	}
	return client
}
