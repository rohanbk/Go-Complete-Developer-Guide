package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	BOT_ID
}
type spanishBot struct {
	BOT_ID
}

type BOT_ID struct {
	id int
}

func main() {
	eb := englishBot{
		BOT_ID{
			id: 1,
		},
	}
	sb := spanishBot{
		BOT_ID{
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
