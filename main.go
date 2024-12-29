package main

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/getlantern/systray"
	"github.com/gorilla/websocket"
)

//go:embed assets/icons/0-1.ico
var icon0 []byte

//go:embed assets/icons/1-1-tr.ico
var icon1 []byte

//go:embed assets/icons/2-1-tr.ico
var icon2 []byte

//go:embed assets/icons/3-1-tr.ico
var icon3 []byte

//go:embed assets/icons/4-1-tr.ico
var icon4 []byte

//go:embed assets/icons/5-1-tr.ico
var icon5 []byte

//go:embed assets/icons/6-1-tr.ico
var icon6 []byte

//go:embed assets/icons/7-1-tr.ico
var icon7 []byte

//go:embed assets/icons/8-1-tr.ico
var icon8 []byte

//go:embed assets/icons/9-1-tr.ico
var icon9 []byte

var workspacesById = make(map[string]string)
var iconsByNumber = make(map[string][]byte)
var ready = false

func main() {
	log.Println("Starting GlazeWmSysTrayIndicator")

	err := loadIcons()
	if err != nil {
		log.Fatal(err)
	}

	connection, _, err := websocket.DefaultDialer.Dial("ws://localhost:6123", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = loadWorkspaces(connection)
	if err != nil {
		log.Fatal(err)
	}

	go reactToWorkspaceChanges(connection)
	systray.Run(onReady, onExit)
}

func onReady() {
	ready = true
	systray.SetTooltip("GlazeWM indicator")
	systray.AddMenuItem("hello", "world")
	systray.SetIcon(iconsByNumber["0"])
}

func onExit() {
	// clean up here
}

func reactToWorkspaceChanges(connection *websocket.Conn) {
	connection.WriteMessage(1, []byte("sub -e focus_changed workspace_activated"))
	for {
		var message GlazeWmMessage[EventWrapper]
		_, buf, err := connection.ReadMessage()
		if err != nil {
			log.Println("Read error", err)
			continue
		}

		log.Println("Received ", string(buf))

		err = json.Unmarshal(buf, &message)
		if err != nil {
			log.Println("Unmarshal error", err)
			continue
		}

		switch event := message.Data.Value.(type) {
		case FocusChangedEvent:
			onFocusChangedEvent(event)
		case WorkspaceActivatedEvent:
			onWorkspaceActivatedEvent(event)
		default:
			log.Println("Unexpected event type", message.MessageType)
			continue
		}

		// TODO: edge case after registering:
		// 2024/12/09 21:53:04 Received  {"messageType":"client_response","clientMessage":"sub -e focus_changed","data":{"subscriptionId":"08c01082-4b94-4731-b794-1b3147f180c1"},"error":null,"success":true}
	}
}

func onFocusChangedEvent(event FocusChangedEvent) {
	var workspaceId string
	switch s := event.FocusedContainer.Value.(type) {
	case Window:
		workspaceId = s.ParentId
	case Workspace:
		workspaceId = s.Id
	default:
		log.Fatalf("Unknown type %s", s)
	}

	workspaceName, ok := workspacesById[workspaceId]
	if !ok {
		log.Println("Unknown workspace", event)
	}

	onWorkspaceSelected(workspaceName)
}

func onWorkspaceSelected(workspaceName string) {
	log.Println("Switching to workspace '", workspaceName, "'")

	if ready {
		if workspaceName != "" {
			systray.SetIcon(iconsByNumber[workspaceName])
		} else {
			systray.SetIcon(iconsByNumber["0"])
		}
	}
}

func onWorkspaceActivatedEvent(event WorkspaceActivatedEvent) {
	workspacesById[event.ActivatedWorkspace.Id] = event.ActivatedWorkspace.Name
}

func loadIcons() error {
	log.Println("Loading icons")
	for i := 0; i <= 9; i++ {
		// Icon source: https://github.com/urob/komotray
		iconsByNumber["0"] = icon0
		iconsByNumber["1"] = icon1
		iconsByNumber["2"] = icon2
		iconsByNumber["3"] = icon3
		iconsByNumber["4"] = icon4
		iconsByNumber["5"] = icon5
		iconsByNumber["6"] = icon6
		iconsByNumber["7"] = icon7
		iconsByNumber["8"] = icon8
		iconsByNumber["9"] = icon9
	}

	return nil
}

func loadWorkspaces(connection *websocket.Conn) error {
	log.Println("Loading workspaces")

	connection.WriteMessage(1, []byte("query workspaces"))
	var queryResponse GlazeWmMessage[Workspaces]
	_, buf, err := connection.ReadMessage()
	// err := connection.ReadJSON(&queryResponse)
	if err != nil {
		return err
	}

	log.Println("Initially received ", string(buf))

	err = json.Unmarshal(buf, &queryResponse)
	if err != nil {
		return err
	}

	for _, w := range queryResponse.Data.Workspaces {
		if w.HasFocus {
			onWorkspaceSelected(w.Name)
		}
		workspacesById[w.Id] = w.Name
	}

	return nil
}
