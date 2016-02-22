package messages

type MessageHandler func(message Message) error

type Message interface {
	Words() []string
	Message() string
	FullMessage() string
	Respond(string)
}
