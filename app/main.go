package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func main() {
	// Load .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set in .env")
	}

	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid TELEGRAM_CHAT_ID: %v", err)
	}

	apiURL := "https://api.telegram.org/bot" + botToken + "/sendMessage"

	for i := 1; i <= 4; i++ {
		message := fmt.Sprintf("Penipu blog #%d", i)
		err := sendMessage(apiURL, chatID, message)
		if err != nil {
			fmt.Printf("Error sending message '%s': %v\n", message, err)
		} else {
			fmt.Printf("Message sent successfully: '%s'\n", message)
		}
	}
}

func sendMessage(apiURL string, chatID int64, text string) error {
	msg := TelegramMessage{
		ChatID: chatID,
		Text:   text,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
