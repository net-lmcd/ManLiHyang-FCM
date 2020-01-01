package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

/**
Router Setting
*/
func ginRouter() *gin.Engine {
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

	router.GET("/fcm/ping", pong)
	router.GET("/fcm", fcmList) // /fcm?offset=0&count=10
	router.POST("/fcm", fcmSendTo)
	router.POST("/fcm/multicast", fcmMultiCast)
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
	offset := gc.DefaultQuery("offset", "0")
	count := gc.DefaultQuery("count", "10")
	if offset == "" || count == "" {
		gc.JSON(http.StatusBadRequest, createBasicBR())
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
	if err := gc.ShouldBindJSON(&request); err != nil {
		gc.JSON(http.StatusBadRequest, createBR(err))
	}

	from := request.From // string
	to := request.To     // []string
	// TODO sendMessage with goroutine
	log.Print("FROM : ", from, "TO : ", to)

	sendMessage(&request)
	gc.JSON(http.StatusOK, gin.H{"status": "ok"})
}

/**
FCM MultiCast
*/
func fcmMultiCast(gc *gin.Context) {
	var request Request
	if err := gc.ShouldBindJSON(&request); err != nil {
		gc.JSON(http.StatusBadRequest, createBR(err))
	}
	gc.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// validator
func (data *Message) validateData() bool {
	return (data.Body == "" || data.Title == "")
}

func main() {
	runtime.GOMAXPROCS(4)

	log.Print("[FCM SERVER START]")
	router := ginRouter()
	router.Run(":5000")
}
