package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	route := r.Group("/")
	route.GET("/getHello", HandleHello)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 80),
		Handler:        r,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20, // 1M
	}
	fmt.Printf("Server starting at http://127.0.0.1:%d", 80)
	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("Listen: %s\n", err)
	}
}

func HandleHello(c *gin.Context) {
	c.JSON(200, Resp{
		Status:    200,
		Content:   "hello world",
		ErrorMsg:  "",
		Timestamp: time.Now(),
	})
}

type Resp struct {
	Status    int         `json:"status"`
	Content   interface{} `json:"content"`
	ErrorMsg  string      `json:"errorMsg"`
	Timestamp interface{} `json:"timestamp"`
}
