package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	. "github.com/h-yamada/movikuma-bot/config"
	. "github.com/h-yamada/movikuma-bot/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := Conf.Validate(); err == nil {
		router := gin.Default()

		router.GET("/webhook", GetWebHook)
		router.POST("/webhook", PostWebHook)

		router.Run(":9000")
	} else {
		fmt.Println("[config error]", err.Error())
	}
}

func init() {
	var configFile string

	flag.StringVar(&configFile, "config-path", "./movikuma-bot.toml", "config-path")
	flag.Parse()

	if _, err := toml.DecodeFile(configFile, &Conf); err != nil {
		log.Println(err)
	}
	log.Println("FB Token:", Conf.Facebook.Token)
	log.Println("ElasticSearchUri:", Conf.Movikuma.ElasticSearchUri)
	log.Println("DetailPageUri:", Conf.Movikuma.DetailPageUri)
	log.Println("ImageUri:", Conf.Movikuma.ImageUri)

}
