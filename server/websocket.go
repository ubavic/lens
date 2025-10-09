package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/ubavic/lens/view/component"
	"golang.org/x/net/websocket"
)

type Payload struct {
	Id    string
	Value string
}

func handleWs(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	n, err := ws.Read(buf)
	sessionId := string(buf[:n])

	fmt.Println("connected", sessionId)

	for {
		n, err = ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				removeSession(sessionId)
				return
			}

			time.Sleep(100)
			continue
		}

		payload := Payload{}
		json.Unmarshal(buf[:n], &payload)

		action := component.Action{
			Caller: payload.Id,
			Kind:   1,
			Data:   payload.Value,
		}

		handler, ok := sessions[sessionId].handlers[payload.Id]
		if !ok {
			fmt.Println("handler miss", payload.Id)
			continue
		}

		buf2 := bytes.NewBuffer(nil)
		handler(action).Node().Render(buf2)

		msg, _ := json.Marshal(Payload{Id: payload.Id, Value: buf2.String()})

		ws.Write(msg)
	}
}
