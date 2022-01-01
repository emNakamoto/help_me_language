package ui

import (
    "fmt"
)

func PrintCyan(text string) {
	color := "\033[36m"
	fmt.Println(string(color), text, "\033[0m")
}

func PrintPurple(text string) {
	color := "\033[35m"
	fmt.Println(string(color), text, "\033[0m")
}

func PrintBlue(text string) {
	color := "\033[34m"
	fmt.Println(string(color), text, "\033[0m")
}

func PrintYellow(text string) {
	color := "\033[33m"
	fmt.Println(string(color), text, "\033[0m")
}

func PrintGreen(text string) {
	color := "\033[32m"
	fmt.Println(string(color), text, "\033[0m")
}

func PrintRed(text string) {
	color := "\033[31m"
	fmt.Println(string(color), text, "\033[0m")
}

func Print(text string) {
	fmt.Println(text)
}

