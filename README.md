# Text-Based Adventure Game

A simple interactive text-based adventure game written in Go with multiple story paths and endings.

## How to Use

### Prerequisites
- Go 1.21 or later installed on your system

### Installation & Running

1. **Create a new directory for the game:**
   ```bash
   mkdir text-adventure-game
   cd text-adventure-game
   ```

2. **Create the game files:**
   Copy the provided code into three separate files:
   - `main.go`
   - `game.go`
   - `go.mod`

3. **Initialize and run the game:**
   ```bash
   go mod tidy
   go run .
   ```

### Game Features

- **Multiple Story Paths**: Choose different paths that lead to unique adventures
- **Inventory System**: Collect items that can help you later in the game
- **Health System**: Your choices affect your character's health
- **Multiple Endings**: Discover different winning and losing scenarios
- **Interactive Choices**: Make decisions by entering numbers (1, 2, 3, etc.)

### How to Play

1. **Start the game** by running the program
2. **Enter your character's name** when prompted
3. **Read the story descriptions** carefully
4. **Make choices** by entering the corresponding number
5. **Explore different paths** to discover all possible endings
6. **Collect items** that might help you in future encounters

### Game Structure

The game starts in a forest with three initial paths:
- **Left Path**: Leads to a dark cave with treasure and dangers
- **Right Path**: Leads to a river with mystical discoveries
- **Straight Path**: Leads deeper into the forest with magical encounters

### Tips for Success

- Pay attention to item descriptions - they might be useful later
- Some choices require specific items to succeed
- Explore all paths to experience the full story
- Your health matters in certain encounters

### Replaying

To play again with different choices, simply run the program again:
```bash
go run .
```

Enjoy your adventure through the mystical forest!