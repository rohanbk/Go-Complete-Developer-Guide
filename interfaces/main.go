package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	botID
}
type spanishBot struct {
	botID
}

type botID struct {
	id int
}

func main() {
	eb := englishBot{
		botID{
			id: 1,
		},
	}
	sb := spanishBot{
		botID{
			id: 2,
		},
	}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola! Como estas?"
}
