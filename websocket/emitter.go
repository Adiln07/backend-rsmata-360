package websocket

import "encoding/json"

type WSMesage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func Emit(event string, data interface{}) {
	if HubInstance == nil {
		return
	}

	message := WSMesage{
		Event: event,
		Data:  data,
	}

	jsonData, _ := json.Marshal(message)

	HubInstance.Broadcast <- jsonData
}