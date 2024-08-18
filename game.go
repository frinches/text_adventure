package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameState struct {
	CurrentScene string
	Inventory    []string
	Health       int
	Completed    bool
}

type Game struct {
	state *GameState
	scenes map[string]*Scene
}

type Scene struct {
	Description string
	Choices     []Choice
}

type Choice struct {
	Text    string
	Next    string
	Effect  func(*GameState)
}

func NewGame() *Game {
	game := &Game{
		state: &GameState{
			CurrentScene: "start",
			Health:       100,
			Inventory:    []string{},
		},
		scenes: make(map[string]*Scene),
	}
	
	game.initializeScenes()
	return game
}

func (g *Game) initializeScenes() {
	// Start scene
	g.scenes["start"] = &Scene{
		Description: "You stand at the edge of a mysterious forest. Two paths diverge before you.",
		Choices: []Choice{
			{
				Text: "Take the left path - it looks well-traveled",
				Next: "left_path",
			},
			{
				Text: "Take the right path - it's overgrown but intriguing",
				Next: "right_path",
			},
		},
	}
	
	// Left path scenes
	g.scenes["left_path"] = &Scene{
		Description: "You follow the left path and find an ancient stone bridge crossing a river.",
		Choices: []Choice{
			{
				Text: "Cross the bridge carefully",
				Next: "bridge_crossing",
			},
			{
				Text: "Search the riverbank for items",
				Next: "river_search",
			},
		},
	}
	
	g.scenes["bridge_crossing"] = &Scene{
		Description: "As you cross the bridge, you notice it's unstable. Suddenly, a plank breaks!",
		Choices: []Choice{
			{
				Text: "Jump to the other side",
				Next: "bridge_success",
				Effect: func(gs *GameState) {
					gs.Health -= 10
					fmt.Println("You made it but took 10 damage!")
				},
			},
			{
				Text: "Carefully navigate around the broken plank",
				Next: "bridge_success",
			},
		},
	}
	
	g.scenes["bridge_success"] = &Scene{
		Description: "You successfully cross the bridge and find a friendly village ahead.",
		Choices: []Choice{
			{
				Text: "Enter the village",
				Next: "village",
			},
		},
	}
	
	g.scenes["river_search"] = &Scene{
		Description: "You search the riverbank and find a shiny key half-buried in the mud.",
		Choices: []Choice{
			{
				Text: "Take the key and continue",
				Next: "bridge_crossing",
				Effect: func(gs *GameState) {
					gs.Inventory = append(gs.Inventory, "rusty_key")
					fmt.Println("You added a rusty key to your inventory!")
				},
			},
		},
	}
	
	// Right path scenes
	g.scenes["right_path"] = &Scene{
		Description: "The overgrown path leads to a dark cave entrance. Strange sounds echo from within.",
		Choices: []Choice{
			{
				Text: "Enter the cave",
				Next: "cave_entrance",
			},
			{
				Text: "Search around the cave entrance",
				Next: "cave_search",
			},
		},
	}
	
	g.scenes["cave_entrance"] = &Scene{
		Description: "Inside the cave, you find two tunnels. One glows faintly, the other is pitch black.",
		Choices: []Choice{
			{
				Text: "Take the glowing tunnel",
				Next: "glowing_tunnel",
			},
			{
				Text: "Take the dark tunnel",
				Next: "dark_tunnel",
			},
		},
	}
	
	g.scenes["glowing_tunnel"] = &Scene{
		Description: "The glowing tunnel leads to a chamber filled with magical crystals. A treasure chest sits in the center.",
		Choices: []Choice{
			{
				Text: "Open the chest",
				Next: "treasure_chest",
			},
			{
				Text: "Take some crystals and leave",
				Next: "crystal_escape",
				Effect: func(gs *GameState) {
					gs.Inventory = append(gs.Inventory, "magic_crystal")
					fmt.Println("You took a magic crystal!")
				},
			},
		},
	}
	
	g.scenes["treasure_chest"] = &Scene{
		Description: "The chest is locked. You need a key to open it.",
		Choices: []Choice{
			{
				Text: "Try to open it anyway",
				Next: "chest_failed",
			},
			{
				Text: "Leave the chest",
				Next: "crystal_escape",
			},
		},
	}
	
	g.scenes["chest_failed"] = &Scene{
		Description: "The chest won't budge. It's securely locked.",
		Choices: []Choice{
			{
				Text: "Continue exploring",
				Next: "crystal_escape",
			},
		},
	}
	
	g.scenes["crystal_escape"] = &Scene{
		Description: "You escape the cave with your findings and find yourself back at the forest edge, wiser and richer.",
		Choices: []Choice{
			{
				Text: "End your adventure",
				Next: "end",
				Effect: func(gs *GameState) {
					gs.Completed = true
				},
			},
		},
	}
	
	g.scenes["dark_tunnel"] = &Scene{
		Description: "The dark tunnel leads to a dead end. Bats suddenly swarm around you!",
		Choices: []Choice{
			{
				Text: "Fight off the bats",
				Next: "bat_fight",
				Effect: func(gs *GameState) {
					gs.Health -= 20
					fmt.Println("You fought bravely but took 20 damage from the bats!")
				},
			},
			{
				Text: "Run back to the entrance",
				Next: "cave_entrance",
			},
		},
	}
	
	g.scenes["bat_fight"] = &Scene{
		Description: "After fighting the bats, you find a small alcove with an old map.",
		Choices: []Choice{
			{
				Text: "Take the map and continue",
				Next: "crystal_escape",
				Effect: func(gs *GameState) {
					gs.Inventory = append(gs.Inventory, "old_map")
					fmt.Println("You found an old map!")
				},
			},
		},
	}
	
	g.scenes["cave_search"] = &Scene{
		Description: "Around the cave entrance, you find some edible berries and a sturdy stick.",
		Choices: []Choice{
			{
				Text: "Take the berries and stick, then enter the cave",
				Next: "cave_entrance",
				Effect: func(gs *GameState) {
					gs.Inventory = append(gs.Inventory, "berries", "sturdy_stick")
					gs.Health += 15
					fmt.Println("You ate some berries (+15 health) and took a sturdy stick!")
				},
			},
		},
	}
	
	g.scenes["village"] = &Scene{
		Description: "The villagers welcome you! They offer you shelter and tell you about the forest's secrets.",
		Choices: []Choice{
			{
				Text: "Stay in the village and end your adventure",
				Next: "end",
				Effect: func(gs *GameState) {
					gs.Completed = true
				},
			},
		},
	}
	
	// End scene
	g.scenes["end"] = &Scene{
		Description: "Your adventure comes to an end. You reflect on your journey...",
		Choices: []Choice{
			{
				Text: "Play again",
				Next: "start",
				Effect: func(gs *GameState) {
					// Reset game state
					gs.CurrentScene = "start"
					gs.Inventory = []string{}
					gs.Health = 100
					gs.Completed = false
				},
			},
			{
				Text: "Quit game",
				Next: "quit",
			},
		},
	}
	
	g.scenes["quit"] = &Scene{
		Description: "Thanks for playing!",
		Choices:     []Choice{},
	}
}

func (g *Game) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for !g.state.Completed && g.state.CurrentScene != "quit" {
		currentScene := g.scenes[g.state.CurrentScene]
		
		// Display current status
		fmt.Printf("\n--- Health: %d | Inventory: %v ---\n", g.state.Health, g.state.Inventory)
		fmt.Println(currentScene.Description)
		fmt.Println("\nWhat do you do?")
		
		// Display choices
		for i, choice := range currentScene.Choices {
			fmt.Printf("%d. %s\n", i+1, choice.Text)
		}
		
		// Get player input
		fmt.Print("\nEnter your choice (1-", len(currentScene.Choices), "): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		
		// Validate input
		var choiceIndex int
		_, err := fmt.Sscanf(input, "%d", &choiceIndex)
		if err != nil || choiceIndex < 1 || choiceIndex > len(currentScene.Choices) {
			fmt.Println("Invalid choice! Please enter a number between 1 and", len(currentScene.Choices))
			continue
		}
		
		// Execute choice
		selectedChoice := currentScene.Choices[choiceIndex-1]
		if selectedChoice.Effect != nil {
			selectedChoice.Effect(g.state)
		}
		g.state.CurrentScene = selectedChoice.Next
		
		// Check for game over
		if g.state.Health <= 0 {
			fmt.Println("\n💀 Your health has reached zero! Game Over.")
			break
		}
	}
	
	fmt.Println("\nGame ended. Thanks for playing!")
}