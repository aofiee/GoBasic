package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#0000ff",
		"black": "#000000",
		"grey":  "#00cc00",
	}
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("%v %v", key, value)
	}
}
