package triggers

import (
	"at-terminal-chat-server/models"
	"encoding/json"
	"fmt"
	atClientLib "github.com/AlekseiKromski/at-socket-server/client"
	"github.com/AlekseiKromski/at-socket-server/core"
)

type ToAll struct {
	Client  *core.Client
	Clients []*core.Client
	Data    string
}

func (ta *ToAll) SetClient(client *core.Client) {
	ta.Client = client
}

func (ta *ToAll) SetClients(client []*core.Client) {
	ta.Clients = client
}

func (ta *ToAll) SetData(data string) {
	ta.Data = data
}

func (ta *ToAll) Do() {
	for _, client := range ta.Clients {

		//Ignore the user, who sent this message
		if client.ID == ta.Client.ID {
			continue
		}

		message := models.Message{
			From:    string(client.ID),
			Message: ta.Data,
		}
		encodedMessage, err := json.Marshal(message)
		if err != nil {
			fmt.Println("cannot marshal message")
			return
		}

		response := atClientLib.AtServerResponse{
			ClientActionType: "get_message",
			Data:             string(encodedMessage),
		}
		encoded, err := json.Marshal(response)
		if err != nil {
			fmt.Errorf("error in toJson method: %v", err)
			return
		}
		client.Conn.WriteMessage(1, encoded)
	}
}
