package actions

import (
	"fmt"
	"github.com/AlekseiKromski/at-socket-server/core"
)

type SendMessage struct {
	Data   string
	client *core.Client
}

func (sm *SendMessage) SetData(data string) {
	sm.Data = data
}

func (sm *SendMessage) Do() {
	sm.run()
}
func (sm *SendMessage) TrigType() string {
	return "to_all"
}
func (sm *SendMessage) SetClient(client *core.Client) {
	sm.client = client
}

func (sm *SendMessage) run() {
	fmt.Printf("message was received, will send to other [SERVER:/%s]", sm.client.ID)
}
