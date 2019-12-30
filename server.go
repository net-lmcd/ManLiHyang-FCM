package main

import (
	"fmt"
	"log"
	"net/http"
)

type fcmHTTPHandler struct {
	http.Handler
}

func (handelr *fcmHTTPHandler) ServHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := req.GetBody()
	fmt.Println("BODY : ", body)
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	w.Write([]byte("SUCCESS"))
}

func startServer() {
	http.Handle("/fcm", new(fcmHTTPHandler))
	http.ListenAndServe(":5000", nil)
}

func main() {
	firebaseApp := initializeAppWithServiceAccount()
	if firebaseApp != nil {
		log.Print("[FCM SEND MESSAGE]")
		isComplete := sendToken(firebaseApp)
		if !isComplete {
			log.Fatal("[FCM SEND MESSAGE FAILED, EXIT NOW")
		}
	}

	log.Print("[FCM SERVER START]")
	startServer()
}
