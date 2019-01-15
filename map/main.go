package main

import "fmt"

func main() {

	colors := make(map[string]string)
	colors["red"] = "#FFF000"
	addColor(colors, "white", "#000000")
	printColors(colors)

}

func addColor(colors map[string]string, color string, hex string) {
	colors[color] = hex
}

func deleteColor(colors map[string]string, color string) {
	delete(colors, color)
}

func printColors(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("%v, %v\n", color, hex)
	}
}
