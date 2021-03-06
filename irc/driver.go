package irc

import (
	"bufio"
	"fmt"
	"github.com/hagbarddenstore/kodapor-bot/messages"
	"net"
	"os"
	"strings"
	"time"
)

// Driver defines the IRC driver.
type Driver struct {
	host                    string
	port                    int
	username                string
	channels                []string
	connection              net.Conn
	write                   chan string
	messageReceivedHandlers []messages.MessageHandler
}

// New creates a new IRC driver.
func New(
	host string,
	port int,
	username string,
	channels []string) (*Driver, error) {
	driver := &Driver{
		host:     host,
		port:     port,
		username: username,
		channels: channels,
		write:    make(chan string, 10),
	}

	return driver, nil
}

// Connect the driver to the IRC host and start the message handler.
func (d *Driver) Connect() {
	address := fmt.Sprintf("%s:%d", d.host, d.port)

	connection, err := net.DialTimeout("tcp", address, 30*time.Second)

	if err != nil {
		fmt.Fprintf(os.Stderr, "irc: %s\n", err.Error())

		return
	}

	d.connection = connection

	go d.startReadLoop()
	go d.startWriteLoop()

	d.write <- fmt.Sprintf("NICK %s\r\n", d.username)
	d.write <- fmt.Sprintf("USER %s 127.0.0.1 127.0.0.1 :%s\r\n", d.username, d.username)

	for _, channel := range d.channels {
		d.write <- fmt.Sprintf("JOIN %s\r\n", channel)

		d.write <- fmt.Sprintf("PRIVMSG %s :%s\r\n", channel, "Never fear, I is here")
	}
}

func (d *Driver) Disconnect() {
	// ???
}

func (d *Driver) MessageReceived(handler messages.MessageHandler) {
	d.messageReceivedHandlers = append(d.messageReceivedHandlers, handler)
}

func (d *Driver) Write(message string) {
	d.write <- message
}

func (d *Driver) startReadLoop() {
	buffer := bufio.NewReaderSize(d.connection, 1024)

	for {
		fullMessage, err := buffer.ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "irc: %s\n", err.Error())

			continue
		}

		fmt.Fprintf(os.Stdout, "irc: received message %s\n", strings.TrimSpace(fullMessage))

		message := NewMessage(d, fullMessage)

		fmt.Fprintf(os.Stdout, "irc: parsed message %s\n", message.String())

		if len(d.messageReceivedHandlers) > 0 {
			for _, handler := range d.messageReceivedHandlers {
				handler(message)
			}
		}
	}
}

func (d *Driver) startWriteLoop() {
	for {
		select {
		case message, ok := <-d.write:
			if !ok || message == "" {
				return
			}

			fmt.Fprintf(os.Stdout, "irc: sending message %s\n", strings.TrimSpace(message))

			_, err := d.connection.Write([]byte(message))

			if err != nil {
				fmt.Fprintf(os.Stderr, "irc: error sending message %s\n", err.Error())

				return
			}
		}
	}
}
