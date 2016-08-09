package model

import (
	"log"

	. "github.com/h-yamada/movikuma-bot/config"
	"github.com/h-yamada/tdlog"
)

type SendData struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func SendLogToYBI(sender string, message string) error {
	sendData := SendData{Sender: sender, Message: message}

	t := tdlog.NewTDLog(Conf.YBI.Endpoint, Conf.YBI.ApiKey)
	if err := t.SendLog(sendData); err != nil {
		log.Println(err)
		return nil
	}
	return nil
}
