package main

import (
	"fmt"
)

func main() {
	var message, howIGreetingWord string
	howIGreetingWord = `Today I want to say simple "yes".`
	setMessage(&message, howIGreetingWord)

	fmt.Println(message)
}

func setMessage(messageVariable *string, message string) {
	*messageVariable = message
}
