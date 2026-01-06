package main

import "fmt"

func emailWorker(id int, recipientChan chan Recipient, done chan bool) {

	for recipient := range recipientChan {
		fmt.Println(id, recipient)
	}

	done <- true
}
