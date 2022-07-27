package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

	"github.com/VictorManduca/websocket-golang/server/config"
)

func main() {
	server := gin.Default()
	websocket := melody.New()

	server.GET("/ws", func(c *gin.Context) {
		websocket.HandleRequest(c.Writer, c.Request)
	})

	websocket.HandleMessage(func(s *melody.Session, msg []byte) {
		websocket.Broadcast(msg)
	})

	server.Run(":" + config.PORT)
}
