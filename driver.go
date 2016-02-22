package main

import (
	"github.com/hagbarddenstore/kodapor-bot/messages"
)

type Driver interface {
	Connect()
	Disconnect()
	MessageReceived(handler messages.MessageHandler)
}
