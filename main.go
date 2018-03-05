//go:generate swagger generate spec

package main

import (
	logHook "github.com/b6pzeusbc54tvhw5jgpyw8pwz2x6gs/microservice-k8s-study-notice/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/b6pzeusbc54tvhw5jgpyw8pwz2x6gs/microservice-k8s-study-notice/controllers"
)

func main() {
	logHook.Init()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		/*** START doNotDisturb ***/
		globalNotiSetting := new(controllers.GlobalNotiSetting)
		// swagger:route GET /pets pets users listPets
		//
		// Lists pets filtered by some parameters.
		//
		// This will show all available pets by default.
		// You can get the pets that are out of stock
		//
		//     Consumes:
		//     - application/json
		//     - application/x-protobuf
		//
		//     Produces:
		//     - application/json
		//     - application/x-protobuf
		//
		//     Schemes: http, https, ws, wss
		//
		//     Security:
		//       api_key:
		//       oauth: read, write
		//
		//     Responses:
		//       default: genericError
		//       200: someResponse
		//       422: validationError
		v1.GET("/globalSetting/:userId", globalNotiSetting.One)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
	log.Info("listen and serve on 0.0.0.0:8080")
}
