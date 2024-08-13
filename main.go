package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== THE FOREST ADVENTURE ===")
	fmt.Println("Welcome to the enchanted forest!")
	fmt.Println("Your choices will shape your destiny...\n")
	
	// Initialize game state
	game := NewGame()
	
	// Start the game
	game.Start()
}