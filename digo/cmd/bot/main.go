package main

import (
	db "digo/internal/database"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	commands = []string{"test", "hots", "music", "recommend"}
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v\n", err)
	}
}

func main() {
	// Create a new Discord session using the token from the .env file
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN must be set in .env file")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// Set required intents
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Register event handlers
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(commandHandler)

	// Open connection to Discord
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	// Connect to database
	if err := connectToDB(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Wait for interrupt signal to gracefully shutdown
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	// Clean up before exiting
	dg.Close()
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("We have logged in as %s\n", s.State.User.Username)

	// Load all command handlers
	for _, cmd := range commands {
		fmt.Printf("Loading %s command module...\n", cmd)
		// Here you would initialize your command handlers
		// Implementation will depend on how you structure your commands
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Log messages to file
	logMessage := fmt.Sprintf("%s > (%s) %s: %s\n",
		m.ChannelID,
		m.Author.ID,
		m.Author.Username,
		m.Content)

	f, err := os.OpenFile("discord_chatting_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(logMessage); err != nil {
		log.Printf("Error writing to log file: %v\n", err)
	}
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if message starts with command prefix
	if len(m.Content) == 0 || m.Content[0] != '!' {
		return
	}

	// Here you would implement your command routing logic
	// This will be expanded based on your specific commands
}

// Import the db package in your imports:
// "yourproject/db"

// Database connection function
func connectToDB() error {
	return db.ConnectToDB()
}

// Add this to your cleanup in main():
// db.DisconnectDB()
