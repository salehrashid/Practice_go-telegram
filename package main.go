package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	botToken = "1234567890:WWCo_arGTsf8gqWudtgCJ40y_wFrDTYuEeo" // Replace with your bot token
	apiURL   = "https://api.telegram.org/bot" + botToken + "/sendMessage"
)

type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func main() {
	chatID := int64(1234567890) // Replace with the actual chat ID

	for i := 1; i <= 4; i++ {
		message := fmt.Sprintf("Penipu blog #%d", i)
		err := sendMessage(chatID, message)
		if err != nil {
			fmt.Printf("Error sending message '%s': %v\n", message, err)
		} else {
			fmt.Printf("Message sent successfully: '%s'\n", message)
		}
	}
}

func sendMessage(chatID int64, text string) error {
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
