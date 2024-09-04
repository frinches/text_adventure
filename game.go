package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	scanner *bufio.Scanner
	player  *Player
}

type Player struct {
	Name  string
	Health int
	Items  []string
}

func NewGame() *Game {
	return &Game{
		scanner: bufio.NewScanner(os.Stdin),
		player: &Player{
			Health: 100,
			Items:  []string{},
		},
	}
}

func (g *Game) Start() {
	fmt.Println("=== WELCOME TO THE FOREST ADVENTURE ===")
	fmt.Print("Enter your name: ")
	g.scanner.Scan()
	g.player.Name = g.scanner.Text()
	
	fmt.Printf("\nHello, %s! Your adventure begins...\n\n", g.player.Name)
	g.forestPath()
}

func (g *Game) getUserInput(prompt string) string {
	fmt.Print(prompt)
	g.scanner.Scan()
	return strings.TrimSpace(g.scanner.Text())
}

func (g *Game) forestPath() {
	fmt.Println("You find yourself at a crossroad in a dense forest.")
	fmt.Println("1. Take the left path - leads to a dark cave")
	fmt.Println("2. Take the right path - leads to a river")
	fmt.Println("3. Go straight ahead - deeper into the forest")
	
	choice := g.getUserInput("Choose your path (1-3): ")
	
	switch choice {
	case "1":
		g.darkCave()
	case "2":
		g.riverPath()
	case "3":
		g.deepForest()
	default:
		fmt.Println("Invalid choice! Please choose 1, 2, or 3.")
		g.forestPath()
	}
}

func (g *Game) darkCave() {
	fmt.Println("\nYou enter a dark, damp cave.")
	fmt.Println("You see two tunnels ahead.")
	fmt.Println("1. Take the narrow tunnel on the left")
	fmt.Println("2. Take the wide tunnel on the right")
	fmt.Println("3. Go back to the forest")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		g.treasureRoom()
	case "2":
		g.dragonEncounter()
	case "3":
		fmt.Println("You return to the forest crossroad.")
		g.forestPath()
	default:
		fmt.Println("Invalid choice!")
		g.darkCave()
	}
}

func (g *Game) riverPath() {
	fmt.Println("\nYou arrive at a fast-flowing river.")
	fmt.Println("A small boat is tied to a nearby tree.")
	fmt.Println("1. Take the boat across the river")
	fmt.Println("2. Follow the river downstream")
	fmt.Println("3. Go back to the forest")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		g.mysteriousIsland()
	case "2":
		g.ancientTemple()
	case "3":
		fmt.Println("You return to the forest crossroad.")
		g.forestPath()
	default:
		fmt.Println("Invalid choice!")
		g.riverPath()
	}
}

func (g *Game) deepForest() {
	fmt.Println("\nYou venture deeper into the forest.")
	fmt.Println("You encounter a friendly elf who offers you help.")
	fmt.Println("1. Accept the elf's help")
	fmt.Println("2. Politely decline and continue alone")
	fmt.Println("3. Ask the elf about the forest")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		g.elfHelp()
	case "2":
		g.lostInForest()
	case "3":
		g.elfInformation()
	default:
		fmt.Println("Invalid choice!")
		g.deepForest()
	}
}

func (g *Game) treasureRoom() {
	fmt.Println("\nYou squeeze through the narrow tunnel and find a hidden treasure room!")
	fmt.Println("Gold coins and jewels sparkle in the dim light.")
	g.player.Items = append(g.player.Items, "Treasure Chest")
	fmt.Println("🎉 YOU FOUND THE TREASURE! YOU WIN! 🎉")
	g.endGame()
}

func (g *Game) dragonEncounter() {
	fmt.Println("\nYou enter a large chamber and encounter a sleeping dragon!")
	fmt.Println("1. Try to sneak past the dragon")
	fmt.Println("2. Attack the dragon")
	fmt.Println("3. Retreat quietly")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		if len(g.player.Items) > 0 {
			fmt.Println("You successfully sneak past the dragon and find the treasure!")
			g.treasureRoom()
		} else {
			fmt.Println("The dragon wakes up! You barely escape with your life.")
			g.player.Health -= 30
			fmt.Printf("Health remaining: %d\n", g.player.Health)
			g.darkCave()
		}
	case "2":
		fmt.Println("You attack the dragon but it's too powerful!")
		fmt.Println("💀 The dragon defeats you. GAME OVER!")
		os.Exit(0)
	case "3":
		fmt.Println("You retreat back to the cave entrance.")
		g.darkCave()
	default:
		fmt.Println("Invalid choice!")
		g.dragonEncounter()
	}
}

func (g *Game) mysteriousIsland() {
	fmt.Println("\nYou cross the river and discover a mysterious island!")
	fmt.Println("An ancient wizard approaches you.")
	fmt.Println("1. Talk to the wizard")
	fmt.Println("2. Explore the island")
	fmt.Println("3. Return to the river")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		g.wizardEncounter()
	case "2":
		fmt.Println("You explore the island and find a magical amulet!")
		g.player.Items = append(g.player.Items, "Magical Amulet")
		fmt.Println("You return to the river.")
		g.riverPath()
	case "3":
		fmt.Println("You return to the river.")
		g.riverPath()
	default:
		fmt.Println("Invalid choice!")
		g.mysteriousIsland()
	}
}

func (g *Game) ancientTemple() {
	fmt.Println("\nFollowing the river, you discover an ancient temple!")
	fmt.Println("The temple doors are sealed with magical runes.")
	fmt.Println("1. Try to decipher the runes")
	fmt.Println("2. Look for another entrance")
	fmt.Println("3. Return to the river")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		if contains(g.player.Items, "Magical Amulet") {
			fmt.Println("The amulet glows and the temple doors open!")
			fmt.Println("Inside, you find ancient knowledge and wisdom.")
			fmt.Println("🧠 YOU GAINED ANCIENT WISDOM! YOU WIN! 🧠")
			g.endGame()
		} else {
			fmt.Println("The runes are too complex to understand without magical help.")
			fmt.Println("You return to the river.")
			g.riverPath()
		}
	case "2":
		fmt.Println("You find a hidden passage but it's blocked by rubble.")
		fmt.Println("You return to the river.")
		g.riverPath()
	case "3":
		fmt.Println("You return to the river.")
		g.riverPath()
	default:
		fmt.Println("Invalid choice!")
		g.ancientTemple()
	}
}

func (g *Game) elfHelp() {
	fmt.Println("\nThe elf gives you a magical compass and a healing potion!")
	g.player.Items = append(g.player.Items, "Magical Compass", "Healing Potion")
	g.player.Health = 100
	fmt.Println("Your health is restored to 100!")
	fmt.Println("The elf guides you to safety.")
	fmt.Println("🌟 YOU COMPLETED YOUR JOURNEY SAFELY! YOU WIN! 🌟")
	g.endGame()
}

func (g *Game) lostInForest() {
	fmt.Println("\nYou continue alone but soon get lost in the dense forest.")
	fmt.Println("Days pass as you wander without direction...")
	fmt.Println("💀 You succumb to hunger and exhaustion. GAME OVER!")
	os.Exit(0)
}

func (g *Game) elfInformation() {
	fmt.Println("\nThe elf tells you about a hidden treasure in the dark cave.")
	fmt.Println("He also warns you about a dragon guarding it.")
	fmt.Println("You thank the elf and continue your journey.")
	g.player.Items = append(g.player.Items, "Cave Map")
	g.deepForest()
}

func (g *Game) wizardEncounter() {
	fmt.Println("\nThe wizard offers to teach you magic in exchange for help.")
	fmt.Println("1. Accept the wizard's offer")
	fmt.Println("2. Decline and explore more")
	fmt.Println("3. Ask about the island's secrets")
	
	choice := g.getUserInput("What do you do? (1-3): ")
	
	switch choice {
	case "1":
		fmt.Println("You spend years learning magic from the wizard.")
		fmt.Println("🔮 YOU BECOME A POWERFUL WIZARD! YOU WIN! 🔮")
		g.endGame()
	case "2":
		fmt.Println("You decline and continue exploring the island.")
		g.mysteriousIsland()
	case "3":
		fmt.Println("The wizard reveals the island's ancient magic but warns you to leave.")
		fmt.Println("You take his advice and return to the river.")
		g.riverPath()
	default:
		fmt.Println("Invalid choice!")
		g.wizardEncounter()
	}
}

func (g *Game) endGame() {
	fmt.Printf("\n=== GAME COMPLETED ===\n")
	fmt.Printf("Player: %s\n", g.player.Name)
	fmt.Printf("Final Health: %d\n", g.player.Health)
	fmt.Printf("Items Collected: %v\n", g.player.Items)
	fmt.Println("Thanks for playing!")
	os.Exit(0)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}