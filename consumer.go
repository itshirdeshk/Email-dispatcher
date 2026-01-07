package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, recipientChan chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range recipientChan {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "Just testing our email campaign.")

		// msg := []byte(formattedMsg)

		msg, err := executeTempplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d, parsing email template for %s", id, recipient.Email)
			continue
		}

		fmt.Printf("Worker %d: Sending email to %s\n", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "hirdesh@gmail.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(nil)
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Sent email to %s\n", id, recipient.Email)
	}
}
