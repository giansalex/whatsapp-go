package cl

import (
	"time"

	"github.com/Rhymen/go-whatsapp"
)

// Message content body whatsapp text message
type Message struct {
	ID string

	From string

	Text string

	Time time.Time

	Source whatsapp.TextMessage
}
