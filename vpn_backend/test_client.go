// test_client.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Create a new VPN session
	sessionID := createSession()
	fmt.Printf("Created session: %s\n", sessionID)

	// Connect to WebSocket
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	queryParams := url.Values{}
	queryParams.Add("peerID", "testPeer1")
	u.RawQuery = queryParams.Encode()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// Send test message
	message := map[string]interface{}{
		"type":      "test",
		"sessionId": sessionID,
		"payload":   "Hello VPN Server!",
	}

	messageBytes, _ := json.Marshal(message)
	err = c.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// Wait for interrupt
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func createSession() string {
	url := "http://localhost:8080/create-session"
	jsonStr := []byte(`{}`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["session_id"].(string)
}
