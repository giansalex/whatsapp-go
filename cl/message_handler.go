package cl

import (
	"fmt"
	"os"

	"github.com/Rhymen/go-whatsapp"
)

type messageListener func(Message)

type messageHandler struct {
	listener messageListener
}

func (h *messageHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
}

func (h *messageHandler) HandleTextMessage(message whatsapp.TextMessage) {
	if message.Info.FromMe {
		return
	}

	mapMessage := Message{
		From: message.Info.RemoteJid,
		Text: message.Text,
	}

	h.listener(mapMessage)
}
