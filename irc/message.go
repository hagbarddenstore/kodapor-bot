package irc

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	driver      *Driver
	fullMessage string
	source      string
	command     string
	target      string
	message     string
	words       []string
}

func NewMessage(driver *Driver, fullMessage string) *Message {
	fullMessage = strings.TrimSuffix(fullMessage, "\r\n")

	if len(fullMessage) < 5 {
		fmt.Fprintln(os.Stderr, "irc: received an invalid message")

		return nil
	}

	message := &Message{
		driver:      driver,
		fullMessage: fullMessage,
	}

	if fullMessage[0] == ':' {
		if i := strings.Index(fullMessage, " "); i > -1 {
			message.source = fullMessage[1:i]

			fullMessage = fullMessage[i+1:]
		}
	}

	parts := strings.SplitN(fullMessage, " :", 2)

	if len(parts) == 2 {
		message.message = parts[1]

		message.words = strings.Split(parts[1], " ")
	}

	arguments := strings.Split(parts[0], " ")

	if len(arguments) == 2 {
		message.command = arguments[0]
		message.target = arguments[1]
	}

	return message
}

func (m *Message) Words() []string {
	return m.words
}

func (m *Message) Message() string {
	return m.message
}

func (m *Message) FullMessage() string {
	return m.fullMessage
}

func (m *Message) Respond(message string) {
	m.driver.Write(fmt.Sprintf("PRIVMSG %s :%s\r\n", m.target, message))
}

func (m *Message) String() string {
	properties := map[string]interface{}{
		"full_message": m.fullMessage,
		"source":       m.source,
		"command":      m.command,
		"target":       m.target,
		"message":      m.message,
		"words":        m.words,
	}

	output, _ := json.Marshal(properties)

	return string(output)
}
