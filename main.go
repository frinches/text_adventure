package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== THE FORGOTTEN TEMPLE ADVENTURE ===")
	fmt.Println("Welcome, brave adventurer!")
	fmt.Println("Your quest begins now...\n")
	
	game := NewGame()
	game.Start()
}