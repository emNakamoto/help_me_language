package main

import (
    "fmt"
)

func printCyan(text string) {
	color := "\033[36m"
	fmt.Println(string(color), text, "\033[0m")
}

func printPurple(text string) {
	color := "\033[35m"
	fmt.Println(string(color), text, "\033[0m")
}

func printBlue(text string) {
	color := "\033[34m"
	fmt.Println(string(color), text, "\033[0m")
}

func printYellow(text string) {
	color := "\033[33m"
	fmt.Println(string(color), text, "\033[0m")
}

func printGreen(text string) {
	color := "\033[32m"
	fmt.Println(string(color), text, "\033[0m")
}

func printRed(text string) {
	color := "\033[31m"
	fmt.Println(string(color), text, "\033[0m")
}

func print(text string) {
	fmt.Println(text)
}
