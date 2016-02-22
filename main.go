package main

import (
	"fmt"
	"github.com/hagbarddenstore/kodapor-bot/irc"
	"github.com/hagbarddenstore/kodapor-bot/version"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Starting %s %s\n", version.Name, version.Version)

	driverName := strings.ToLower(os.Getenv("KAB_DRIVER"))

	fmt.Fprintf(os.Stdout, "main: using driver %s\n", driverName)

	switch driverName {
	case "irc":
		host := os.Getenv("KAB_IRC_HOST")
		username := os.Getenv("KAB_IRC_USERNAME")
		channels := strings.Split(os.Getenv("KAB_IRC_CHANNELS"), ",")

		port, err := strconv.Atoi(os.Getenv("KAB_IRC_PORT"))

		if err != nil {
			fmt.Fprintf(os.Stdout, "main: %s\n", err.Error())

			return
		}

		driver, err := irc.New(host, port, username, channels)

		if err != nil {
			fmt.Fprintf(os.Stdout, "main: %s\n", err.Error())

			return
		}

		driver.Connect()

	default:
		fmt.Fprintln(os.Stderr, "main: no suitable driver found")

		return
	}

	waitForShutdown()
}

func waitForShutdown() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c

		fmt.Fprintf(os.Stderr, "Quitting!\n")

		os.Exit(0)
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
