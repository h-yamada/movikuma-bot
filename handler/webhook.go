package handler

import (
	"log"
	"net/http"

	. "github.com/h-yamada/movikuma-bot/config"

	"github.com/gin-gonic/gin"
)

func GetWebHook(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	log.Println("mode=", mode)
	log.Println("token=", token)
	if mode == "subscribe" && token == Conf.Facebook.Token {
		challenge := c.Query("hub.challenge")
		log.Println("challenge=", challenge)
		c.String(http.StatusOK, challenge)
	} else {
		c.String(http.StatusNotFound, "token error")
	}
}
