package main

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	fmt.Println("Email Dispatcher")
	recipientChannel := make(chan Recipient)

	var wg sync.WaitGroup

	go loadRecipient("users.csv", recipientChannel)

	emailWorkerCount := 5

	for i := 1; i <= emailWorkerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()

}

func executeTempplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tmplBuf bytes.Buffer
	err = t.Execute(&tmplBuf, r)
	if err != nil {
		return "", err
	}
	return tmplBuf.String(), nil
}
