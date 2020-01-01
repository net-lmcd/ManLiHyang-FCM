package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

/**
Router Setting
*/
func ginRouter(app *firebase.App) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.GET("/ping", pong)
	router.GET("/fcm/:offset/:count", fcmList)
	router.POST("fcm", fcmSendTo)
	router.POST("/fcm/multicast", fcmMultiCast)
	router.Run(":5000")
	return router
}

//pong
func pong(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"response": "pong"})
}

/**
GET FCM Token List, offset / count
*/
func fcmList(gc *gin.Context) {
	offset := gc.Params.ByName("offset")
	count := gc.Params.ByName("count")
	if offset == "" || count == "" {
		gc.JSON(http.StatusBadRequest, createBR())
	}
	fmt.Print("offset : ", offset, "count : ", count)

	// io

	slice := Slices{[]string{"a", "b", "c", "d", "e"}}
	gc.JSON(http.StatusOK, slice)
}

/**
POST FCM, Send Message From A To B
*/
func fcmSendTo(gc *gin.Context) {
	var request Request
	gc.BindJSON(&request)
	if request.Code == 0 || request.From == "" || request.To == nil || request.Message.validateData() {
		gc.JSON(http.StatusBadRequest, createBR())
	}

	from := request.From // string
	to := request.To     // []string
	message := creatMessage(request.Message.Title, request.Message.Body)
	// TODO sendMessage with goroutine
	log.Print("FROM : ", from, "TO : ", to, "MESSSAGE : ", message)
	gc.JSON(http.StatusOK, gin.H{"status": "ok"})
}

/**
FCM MultiCast
*/
func fcmMultiCast(gc *gin.Context) {
	var request Request
	gc.BindJSON(&request)
	if request.Code == 0 || request.To == nil || request.Message.validateData() {
		gc.JSON(http.StatusBadRequest, createBR())
		return
	}
	gc.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// validator
func (data *Message) validateData() bool {
	return (data.Body == "" || data.Title == "")
}

func main() {
	firebaseApp := initializeAppWithServiceAccount() // TODO to Singleton
	if firebaseApp != nil {
		log.Print("[FCM SEND MESSAGE]")
		isComplete := sendToken(firebaseApp)
		if !isComplete {
			log.Fatal("[FCM SEND MESSAGE FAILED, EXIT NOW")
		}
	}

	log.Print("[FCM SERVER START]")
}
