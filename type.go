package main

import "github.com/gin-gonic/gin"

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
	S []string `json:"fcm_tokens"`
}

//Message for fcm message
type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

//Request object format, For MultiCast / 운영용
type Request struct {
	Code    int      `json:"service_code"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Message Message  `json:"data"`
}

func createBR() gin.H {
	return gin.H{
		"status": "BAD REQUEST",
		"error":  "check your request",
	}
}
