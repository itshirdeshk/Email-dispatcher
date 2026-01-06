package main

import "fmt"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	fmt.Println("Email Dispatcher")
	recipientChannel := make(chan Recipient)

	done := make(chan bool)

	go loadRecipient("users.csv", recipientChannel)

	go emailWorker(1, recipientChannel, done)

	<-done
}
