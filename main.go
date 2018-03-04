package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	logHook "github.com/b6pzeusbc54tvhw5jgpyw8pwz2x6gs/microservice-k8s-study-notice/logger"
	"github.com/b6pzeusbc54tvhw5jgpyw8pwz2x6gs/microservice-k8s-study-notice/router"
)

func main() {
	logHook.Init()

	log.Info("hi")
	r := gin.Default()
	r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

	}

	v1 :=
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	r.Run() // listen and serve on 0.0.0.0:8080
}
