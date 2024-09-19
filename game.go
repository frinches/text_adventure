package main

import (
	"fmt"
	"strings"
)

type Game struct {
	currentScene *Scene
	player       *Player
}

type Player struct {
	name  string
	items []string
}

func NewGame() *Game {
	player := &Player{name: "Adventurer"}
	game := &Game{player: player}
	game.setupScenes()
	return game
}

func (g *Game) setupScenes() {
	// Create all scenes
	intro := &Scene{
		id:          "intro",
		description: "You stand before an ancient temple entrance. Two paths diverge before you.",
		choices: []Choice{
			{text: "Take the left path into the dark forest", nextSceneID: "forest"},
			{text: "Take the right path up the mountain trail", nextSceneID: "mountain"},
		},
	}

	forest := &Scene{
		id:          "forest",
		description: "The forest is dense and mysterious. You hear strange whispers in the wind.",
		choices: []Choice{
			{text: "Follow the whispers deeper into the forest", nextSceneID: "whispers"},
			{text: "Search for a hidden path", nextSceneID: "hidden_path"},
			{text: "Return to the temple entrance", nextSceneID: "intro"},
		},
	}

	mountain := &Scene{
		id:          "mountain",
		description: "The mountain trail is steep but offers a breathtaking view. You see a cave entrance.",
		choices: []Choice{
			{text: "Enter the cave", nextSceneID: "cave"},
			{text: "Continue climbing the mountain", nextSceneID: "summit"},
			{text: "Return to the temple entrance", nextSceneID: "intro"},
		},
	}

	whispers := &Scene{
		id:          "whispers",
		description: "You discover an ancient druid who offers you wisdom. 'Choose your path wisely,' he says.",
		choices: []Choice{
			{text: "Accept the druid's guidance", nextSceneID: "druid_guidance"},
			{text: "Politely decline and continue exploring", nextSceneID: "forest_explore"},
		},
	}

	hidden_path := &Scene{
		id:          "hidden_path",
		description: "You find a hidden path leading to a magical grove with glowing mushrooms.",
		choices: []Choice{
			{text: "Collect the glowing mushrooms", nextSceneID: "mushrooms"},
			{text: "Follow the path further", nextSceneID: "ancient_tree"},
		},
	}

	cave := &Scene{
		id:          "cave",
		description: "The cave is dark and damp. You see ancient carvings on the walls and hear dripping water.",
		choices: []Choice{
			{text: "Study the ancient carvings", nextSceneID: "carvings"},
			{text: "Follow the sound of water", nextSceneID: "underground_river"},
		},
	}

	summit := &Scene{
		id:          "summit",
		description: "You reach the mountain summit! A magnificent eagle awaits you.",
		choices: []Choice{
			{text: "Approach the eagle", nextSceneID: "eagle"},
			{text: "Look for a way down the other side", nextSceneID: "descent"},
		},
	}

	// Endings
	druid_guidance := &Scene{
		id:          "druid_guidance",
		description: "The druid teaches you ancient secrets. You become the guardian of the forest. THE END.",
		isEnding:    true,
	}

	forest_explore := &Scene{
		id:          "forest_explore",
		description: "You discover a peaceful clearing and decide to make it your home. THE END.",
		isEnding:    true,
	}

	mushrooms := &Scene{
		id:          "mushrooms",
		description: "The mushrooms grant you magical powers! You become a forest mage. THE END.",
		isEnding:    true,
	}

	ancient_tree := &Scene{
		id:          "ancient_tree",
		description: "You find the Heartwood Tree and learn the forest's deepest secrets. THE END.",
		isEnding:    true,
	}

	carvings := &Scene{
		id:          "carvings",
		description: "The carvings reveal the temple's history. You become its historian. THE END.",
		isEnding:    true,
	}

	underground_river := &Scene{
		id:          "underground_river",
		description: "You discover a hidden underground city and become its ruler. THE END.",
		isEnding:    true,
	}

	eagle := &Scene{
		id:          "eagle",
		description: "The eagle offers you a ride to a legendary floating city. THE END.",
		isEnding:    true,
	}

	descent := &Scene{
		id:          "descent",
		description: "You discover a valley of eternal spring and found a new civilization. THE END.",
		isEnding:    true,
	}

	// Set starting scene
	g.currentScene = intro
}

func (g *Game) Start() {
	for g.currentScene != nil {
		g.currentScene.Display()
		
		if g.currentScene.isEnding {
			fmt.Println("\nThanks for playing!")
			break
		}
		
		choice := g.getPlayerChoice()
		g.currentScene = g.getSceneByID(choice.nextSceneID)
		fmt.Println()
	}
}

func (g *Game) getPlayerChoice() Choice {
	for {
		fmt.Print("Enter your choice (1-", len(g.currentScene.choices), "): ")
		
		var input int
		_, err := fmt.Scan(&input)
		
		if err != nil || input < 1 || input > len(g.currentScene.choices) {
			fmt.Println("Invalid choice! Please try again.")
			continue
		}
		
		return g.currentScene.choices[input-1]
	}
}

func (g *Game) getSceneByID(id string) *Scene {
	// In a real implementation, you'd have a map of scenes
	// For simplicity, we'll recreate them (in practice, use a scene manager)
	g.setupScenes()
	switch id {
	case "intro":
		return &Scene{id: "intro", description: "You stand before an ancient temple entrance. Two paths diverge before you.", choices: []Choice{
			{text: "Take the left path into the dark forest", nextSceneID: "forest"},
			{text: "Take the right path up the mountain trail", nextSceneID: "mountain"},
		}}
	case "forest":
		return &Scene{id: "forest", description: "The forest is dense and mysterious. You hear strange whispers in the wind.", choices: []Choice{
			{text: "Follow the whispers deeper into the forest", nextSceneID: "whispers"},
			{text: "Search for a hidden path", nextSceneID: "hidden_path"},
			{text: "Return to the temple entrance", nextSceneID: "intro"},
		}}
	// ... and so on for other scenes
	default:
		return &Scene{id: "end", description: "Your adventure continues...", isEnding: true}
	}
}