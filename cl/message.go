package cl

import "time"

// Message content body whatsapp text message
type Message struct {
	ID string

	From string

	Text string

	Time time.Time
}
