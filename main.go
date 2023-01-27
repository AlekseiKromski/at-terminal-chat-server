package main

import (
	"at-terminal-chat-server/actions"
	"at-terminal-chat-server/triggers"
	"fmt"
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
	_, err := core.Start(actionHandlers, triggerHandlers)
	if err != nil {
		fmt.Println(err)
	}
}
