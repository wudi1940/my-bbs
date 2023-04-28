package main

import (
	"my-bbs/cmd/client"
	"net/http"
	"sync"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	Init()

	r := gin.Default()
	r.Use(client.LoggerToFile())

	r.LoadHTMLGlob("web-front/*")
	r.Use(static.ServeRoot("/", "web-front"))

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/login", handler.Login)

	r.GET("/register", handler.Register)
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}

func Init() {
	var once sync.Once

	once.Do(client.InitLogClient)
}
