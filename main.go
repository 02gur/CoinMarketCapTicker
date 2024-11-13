package main

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	// Define the WebSocket URL
	wsURL := "wss://push.coinmarketcap.com/ws?device=web&client_source=home_page"

	// Parse the URL
	u, err := url.Parse(wsURL)
	if err != nil {
		log.Fatal("URL parsing error:", err)
	}

	// Connect to the WebSocket server
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer c.Close()

	// Prepare the subscription message
	subscriptionMessage := map[string]interface{}{
		"method": "RSUBSCRIPTION",
		"params": []string{"main-site@crypto_price_15s@{}@normal", "1,2,825,1027"},
	}
	message, err := json.Marshal(subscriptionMessage)
	if err != nil {
		log.Fatal("JSON encoding error:", err)
	}

	// Send the message
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("Write error:", err)
	}

	// Receive and print messages from the server
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal("Read error:", err)
		}
		log.Printf("Received: %s", msg)
	}
}
