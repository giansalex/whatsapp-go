# WhatsApp Go ![Go](https://github.com/giansalex/whatsapp-go/workflows/Go/badge.svg)
Whatsapp Bot to send and reply text messages.

![WhatsApp Icon](https://cdn.icon-icons.com/icons2/373/PNG/96/Whatsapp_37229.png)

## Install

```
go get github.com/giansalex/whatsapp-go
```

With [dep](https://golang.github.io/dep/)   
```
dep ensure -add github.com/giansalex/whatsapp-go
```

## Usage

```go
package main

import "github.com/giansalex/whatsapp-go/cl"

func main() {
	client := cl.NewClient()

	client.Listen(func(msg cl.Message) {
		if msg.Text == "Hi" {
			client.SendText(msg.From, "Hello from github!")
		}
	})
}
```

After executing `cl.NewClient()` function, app will create a [WhatsApp connection](https://github.com/Rhymen/go-whatsapp). If you are not logged in, it will print a QR code in the terminal. Scan it with your phone and you are ready to go!

> Bot will remember the session so there is no need to authenticate everytime.

