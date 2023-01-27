package main

import (
	"at-terminal-chat-server/actions"
	"at-terminal-chat-server/triggers"
	"encoding/json"
	"fmt"
	atClient "github.com/AlekseiKromski/at-socket-server/client"
	"github.com/AlekseiKromski/at-socket-server/core"
)

var actionHandlers = []*core.ActionHandler{
	{
		ActionType: "send_message",
		Action:     &actions.SendMessage{},
	},
}

var triggerHandlers = []*core.TriggerHandler{
	{
		TriggerType: "to_all",
		Action:      &triggers.ToAll{},
	},
}

func main() {
	app, err := core.Start(actionHandlers, triggerHandlers)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case hook := <-app.Hooks:
				if hook.HookType == core.CLIENT_ADDED {
					for _, client := range app.Clients {
						if client.ID == hook.Data {
							message := atClient.AtServerResponse{
								ClientActionType: "server_info",
								Data: fmt.Sprintf("At-terminal-chat-server-global|Hi, this is first" +
									"server, thatwas created by at-socket-server. Let's chat with other 🧑🏻‍💻 🎉"),
							}
							encoded, err := json.Marshal(message)
							if err != nil {
								fmt.Println("There is trouble with info marshaling")
							}
							client.Conn.WriteMessage(1, encoded)
						}
					}
				}
			}
		}
	}()

	for {
	}
}
