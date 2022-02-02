package client

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
)

const (
	chatPath = "/chat/ws.rtm.start?token="
)

type Message struct {
	Id       string
	UserName string
	Message  string
	SentTime time.Time
}

func Chat(tokens []string, n int) error {
	for _, token := range tokens {
		url := wsEndpoint + chatPath + token
		if err := openConnection(url, n); err != nil {
			return err
		}
	}
	return nil
}

func openConnection(url string, n int) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	defer conn.Close()
	if err != nil {
		return err
	}

	//next steps should be done in concurrent way
	if err := sendTestMessages(conn, n); err != nil {
		return err
	}

	if countMessages(conn) != n {
		return errors.New("not all messages have been sent")
	}

	return nil
}

func sendTestMessages(conn *websocket.Conn, n int) error {
	for i := 1; i <= n; i++ {
		text := fmt.Sprintf("Test message - %v", i)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			return err
		}
	}
	return nil
}

func countMessages(conn *websocket.Conn) (count int) {
	for {
		_, _, err := conn.ReadMessage()
		if websocket.IsCloseError(err, websocket.CloseGoingAway) && err == io.EOF {
			break
		} else if err != nil {
			return
		}
		count++
	}
	return
}
