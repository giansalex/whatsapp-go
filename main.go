package main

import (
	"github.com/giansalex/whatsapp-go/cl"
)

func main() {
	client := cl.NewClient()

	client.Listen(func(msg cl.Message) {
		if msg.Text == "hola" {
			client.SendText(msg.From, "Hello from *github*!")
		}
	})
}
