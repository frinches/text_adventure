## How to Use the Text-Based Adventure Game

### Prerequisites
- Go 1.21 or later installed on your system

### Installation & Running the Game

1. **Create a new directory for the game:**
   ```bash
   mkdir text-adventure
   cd text-adventure
   ```

2. **Save the files:**
   - Save `main.go`, `game.go`, and `go.mod` in the directory

3. **Run the game:**
   ```bash
   go run main.go game.go
   ```

### Game Features

- **Multiple Story Paths**: Choose between different paths that lead to unique adventures
- **Inventory System**: Collect items that may help you later
- **Health System**: Your choices affect your health
- **Branching Narrative**: Every decision matters and leads to different outcomes
- **Replayability**: Multiple endings and paths to discover

### How to Play

1. **Read the descriptions** carefully for each scene
2. **Choose your actions** by entering the corresponding number (1, 2, 3, etc.)
3. **Manage your health** - some choices may damage you
4. **Collect items** - they might be useful later
5. **Explore different paths** to discover all possible endings

### Game Structure

- **main.go**: Entry point of the application
- **game.go**: Contains all game logic, scenes, and state management
- **go.mod**: Go module configuration

### Extending the Game

You can easily add more scenes and choices by:
1. Adding new scenes to the `initializeScenes()` method in `game.go`
2. Creating new choices with different effects and next scenes
3. Adding more complex inventory interactions

### Example Game Flow
- Start at the forest edge
- Choose between left or right path
- Encounter different challenges based on your choices
- Collect items, fight creatures, solve puzzles
- Reach one of multiple possible endings

Enjoy your adventure in the enchanted forest!