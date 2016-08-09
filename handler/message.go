package handler

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/h-yamada/movikuma-bot/config"
	. "github.com/h-yamada/movikuma-bot/model"

	. "github.com/ymd38/facebook/messenger"

	"github.com/gin-gonic/gin"
)

func PostWebHook(c *gin.Context) {

	receiver := &ReceivedMessage{}
	if err := c.BindJSON(&receiver); err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	for _, messaging := range receiver.Entry[0].Messaging {
		log.Printf("senderid=%s message=%", messaging.Sender.ID, messaging.Message.Text)

		movikuma := new(Movikuma)
		movikumaList, _ := movikuma.Search(messaging.Message.Text)

		SendLogToYBI(messaging.Sender.ID, messaging.Message.Text)

		var m interface{}
		fb := NewFacebookMessenger(Conf.Facebook.Token)
		if movikumaList != nil && len(movikumaList) > 0 {
			elements := []Element{}
			for i := 0; i < len(movikumaList); i++ {
				title := fmt.Sprintf("今、人気の「%s」動画はコチラ！", messaging.Message.Text)
				item_url := Conf.Movikuma.DetailPageUri + movikumaList[i].Key
				image_url := Conf.Movikuma.ImageUri + movikumaList[i].Key + "/thumbnail"
				elements = append(elements, Element{Title: title, ItemUrl: item_url, ImageUrl: image_url})
				if i >= 3 {
					break
				}
			}
			gt := NewGenericTemplate(messaging.Sender.ID, elements)
			m = gt
		} else { //not found movie
			res := fmt.Sprintf("うーん、「%s」ではおすすめの動画はないですね。。", messaging.Message.Text)
			m = NewTextMessage(messaging.Sender.ID, res)
		}

		if err := fb.SendMessage(m); err != nil {
			log.Println(err.Error())
		}
	}

	c.String(http.StatusOK, "OK")
}
