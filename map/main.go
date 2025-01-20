package main

import "fmt"

func main() {
	// Method 1
	colors0 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// Method 2
	// var colors map[string]string

	// Method 3
	colors := make(map[string]string)
	colors["white"] = "#ffffff" // use braket syntax to assign value for a key

	fmt.Println(colors)
	// If key does not exist then val1 will have value type's zero value which is empty string
	val1, ok := colors["white"] // If key exist, ok will be true
	fmt.Println(val1, ok)

	val2, ok := colors["invalidKey"] // If key exist, ok will be true
	fmt.Println(val2, ok)

	printMap(colors0)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
