## How to Use the Text-Based Adventure Game

### Prerequisites
- Go 1.16 or higher installed on your system

### Installation & Running

1. **Create a new directory for the game:**
   ```bash
   mkdir text-adventure-game
   cd text-adventure-game
   ```

2. **Create the files:**
   Copy the code above into the respective files:
   - `main.go`
   - `game.go`
   - `scene.go`
   - `go.mod`

3. **Initialize and run the game:**
   ```bash
   go mod init text-adventure-game
   go run .
   ```

### How to Play

1. **Starting the Game:**
   - Run the program and you'll begin at the temple entrance
   - Read the description of your current location

2. **Making Choices:**
   - You'll be presented with numbered options (1, 2, 3, etc.)
   - Type the number of your choice and press Enter
   - Each choice leads to different story paths and endings

3. **Game Features:**
   - Multiple branching story paths
   - 8 different endings
   - Exploration-based gameplay
   - Simple text-based interface

4. **Game Endings:**
   The game has multiple endings including:
   - Becoming a forest guardian
   - Discovering magical powers
   - Ruling hidden cities
   - And more!

### Game Structure

- **main.go**: Entry point of the application
- **game.go**: Contains game logic and scene management
- **scene.go**: Defines scenes and player choices
- **go.mod**: Go module configuration

### Extending the Game

You can easily add more scenes and story paths by:
1. Adding new scenes in the `setupScenes()` method in `game.go`
2. Creating new choices that link to your scenes
3. Adding new ending conditions

### Troubleshooting

- If you get module errors, run `go mod tidy`
- Ensure all files are in the same directory
- Make sure you're using a compatible Go version

Enjoy your adventure in the Forgotten Temple!