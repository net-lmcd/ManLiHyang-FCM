package main

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

/**
{
	"service_code": 1003,
	"tokens" : [ ],
	"data": {
		"title": "타이틀",
		"bdoy": "내용"
	}
}
**/

//Slices which return tokens
type Slices struct {
	S []string `json:"fcm_tokens" binding:"required"`
}

//Message for fcm message
type Message struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

//Request object format, For MultiCast / 운영용
type Request struct {
	Code    int      `json:"service_code" binding:"required"`
	From    string   `json:"from" binding:"required"`
	To      []string `json:"to" binding:"required"`
	Message Message  `json:"data" binding:"required"`
}

//FirebaseApp for singleton
type FirebaseApp struct {
	app firebase.App
}

func createBasicBR() gin.H {
	return gin.H{
		"status": "BAD REQUEST",
	}
}

func createBR(err error) gin.H {
	return gin.H{
		"status": "BAD REQUEST",
		"error":  err.Error(),
	}
}
