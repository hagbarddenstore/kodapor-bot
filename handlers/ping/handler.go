package ping

import (
	"github.com/hagbarddenstore/kodapor-bot/messages"
)

func Handler(message messages.Message) error {
	if len(message.Words()) >= 1 && message.Words()[0] == ".ping" {
		message.Respond("Pong!")
	}

	return nil
}
