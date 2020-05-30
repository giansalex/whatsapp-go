package cl

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

type messageListener func(Message)

type messageHandler struct {
	listener  messageListener
	afterTime int64
}

func (h *messageHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
}

func (h *messageHandler) HandleTextMessage(message whatsapp.TextMessage) {
	if message.Info.FromMe {
		return
	}

	msgTime := int64(message.Info.Timestamp)

	if h.afterTime > msgTime {
		return
	}

	mapMessage := Message{
		ID:     message.Info.Id,
		From:   message.Info.RemoteJid,
		Text:   message.Text,
		Time:   time.Unix(int64(message.Info.Timestamp), 0),
		Source: message,
	}

	h.listener(mapMessage)
}
