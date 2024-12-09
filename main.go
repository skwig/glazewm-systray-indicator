package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/getlantern/systray"
	"github.com/gorilla/websocket"
)

type GlazeWmMessage[T any] struct {
	Success       bool   `json:"successs"`
	MessageType   string `json:"messageType"`
	ClientMessage string `json:"clientMessage"`
	Data          T      `json:"data"`
	Error         string `json:"error"`
}

type Workspaces struct {
	Workspaces []Workspace `json:"workspaces"`
}

type Workspace struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ParentId    string `json:"parentId"`
	HasFocus    bool   `json:"hasFocus"`
	IsDisplayed bool   `json:"isDisplayed"`
	// Note that there are other properties, but they are not relevant for our usecase
}

type FocusChangedEvent struct {
	EventType        string `json:"eventType"`
	FocusedContainer Window `json:"focusedContainer"`
}

type Window struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	ParentId    string `json:"parentId"`
	HasFocus    bool   `json:"hasFocus"`
	IsDisplayed bool   `json:"isDisplayed"`
	// Note that there are other properties, but they are not relevant for our usecase
}

func reactToWorkspaceChanges(connection *websocket.Conn) {
	connection.WriteMessage(1, []byte("sub -e focus_changed"))
	for {
		var message GlazeWmMessage[FocusChangedEvent]
		_, buf, err := connection.ReadMessage()
		// err := connection.read(&message)
		if err != nil {
			log.Println("read:", err)
			continue
		}

		log.Println("Received ", string(buf))

		err = json.Unmarshal(buf, &message)
		if err != nil {
			log.Println("read:", err)
			continue
		}
		// Initial receive for the following edge cases
		// 2024/12/09 21:53:04 Initially received  {"messageType":"client_response","clientMessage":"query workspaces","data":{"workspaces":[{"type":"workspace","id":"0022b3af-9f96-490d-90a5-0c004cfc77bf","name":"6","displayName":"6","parentId":"f0bdd9e0-37b7-4572-b8ac-fa0203dd0886","children":[{"type":"window","id":"dd46c721-5b17-4068-86d1-1839f405ef67","parentId":"0022b3af-9f96-490d-90a5-0c004cfc77bf","hasFocus":false,"tilingSize":null,"width":199,"height":34,"x":-1379,"y":678,"state":{"type":"minimized"},"prevState":null,"displayState":"shown","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":-1379,"top":678,"right":-1180,"bottom":712},"handle":2950636,"title":"#general | 4sgard - Discord","className":"Chrome_WidgetWin_1","processName":"Discord","activeDrag":null},{"type":"window","id":"5a92cff4-c994-41c5-b03a-1a6f11737ea5","parentId":"0022b3af-9f96-490d-90a5-0c004cfc77bf","hasFocus":false,"tilingSize":0.5,"width":1271,"height":1378,"x":-2554,"y":6,"state":{"type":"tiling"},"prevState":null,"displayState":"shown","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":-1913,"top":271,"right":-646,"bottom":1120},"handle":591760,"title":"RuneLite - helmiescape","className":"SunAwtFrame","processName":"RuneLite","activeDrag":null},{"type":"window","id":"f2c19c00-2e41-45d1-a320-eac591abfb07","parentId":"0022b3af-9f96-490d-90a5-0c004cfc77bf","hasFocus":false,"tilingSize":0.5,"width":1271,"height":1378,"x":-1277,"y":6,"state":{"type":"tiling"},"prevState":null,"displayState":"shown","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":-1929,"top":269,"right":-631,"bottom":1121},"handle":592414,"title":"RuneLite - Skwig","className":"SunAwtFrame","processName":"RuneLite","activeDrag":null}],"childFocusOrder":["f2c19c00-2e41-45d1-a320-eac591abfb07","5a92cff4-c994-41c5-b03a-1a6f11737ea5","dd46c721-5b17-4068-86d1-1839f405ef67"],"hasFocus":false,"isDisplayed":true,"width":2548,"height":1378,"x":-2554,"y":6,"tilingDirection":"horizontal"},{"type":"workspace","id":"58be9b2e-6e75-43d8-b3f1-de82e05bdcb7","name":"1","displayName":"1","parentId":"5c3e169d-4f92-4f5b-9dfd-ce58904806d5","children":[{"type":"window","id":"6dc74ae0-10ee-4819-93aa-195051e80c65","parentId":"58be9b2e-6e75-43d8-b3f1-de82e05bdcb7","hasFocus":false,"tilingSize":null,"width":183,"height":34,"x":1189,"y":678,"state":{"type":"minimized"},"prevState":null,"displayState":"shown","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":1189,"top":678,"right":1372,"bottom":712},"handle":133522,"title":"DiscordChatExporter v2.44.0","className":"Avalonia-cf90996f-b815-42eb-8664-2793c1ec1a93","processName":"DiscordChatExporter","activeDrag":null},{"type":"window","id":"2277d065-40df-48f1-a975-5c1b470500f2","parentId":"58be9b2e-6e75-43d8-b3f1-de82e05bdcb7","hasFocus":true,"tilingSize":1.0,"width":2548,"height":1378,"x":6,"y":6,"state":{"type":"tiling"},"prevState":null,"displayState":"shown","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":134,"top":75,"right":2427,"bottom":1315},"handle":591282,"title":"PowerShell","className":"CASCADIA_HOSTING_WINDOW_CLASS","processName":"WindowsTerminal","activeDrag":null}],"childFocusOrder":["2277d065-40df-48f1-a975-5c1b470500f2","6dc74ae0-10ee-4819-93aa-195051e80c65"],"hasFocus":true,"isDisplayed":true,"width":2548,"height":1378,"x":6,"y":6,"tilingDirection":"horizontal"},{"type":"workspace","id":"b5885c58-2659-4e36-94df-b9bea6e9677a","name":"2","displayName":"2","parentId":"5c3e169d-4f92-4f5b-9dfd-ce58904806d5","children":[],"childFocusOrder":[],"hasFocus":false,"isDisplayed":false,"width":2548,"height":1378,"x":6,"y":6,"tilingDirection":"horizontal"},{"type":"workspace","id":"788b3b32-5de6-4bf9-bd0d-7a7c7c8e8069","name":"3","displayName":"3","parentId":"5c3e169d-4f92-4f5b-9dfd-ce58904806d5","children":[],"childFocusOrder":[],"hasFocus":false,"isDisplayed":false,"width":2548,"height":1378,"x":6,"y":6,"tilingDirection":"horizontal"},{"type":"workspace","id":"a58a6ad0-1a0c-4196-81fd-b0d9dd97ed61","name":"8","displayName":"8","parentId":"7bab97e0-c9c4-4ac9-a479-bfb1c764a53f","children":[{"type":"window","id":"658617bd-b61f-46b2-869d-c12aed35c1f7","parentId":"a58a6ad0-1a0c-4196-81fd-b0d9dd97ed61","hasFocus":false,"tilingSize":0.5,"width":1271,"height":1378,"x":2566,"y":6,"state":{"type":"tiling"},"prevState":null,"displayState":"hidden","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":2694,"top":75,"right":4987,"bottom":1315},"handle":264760,"title":"types - How to convert byte array to string in Go- Stack Overflow - Google Chrome","className":"Chrome_WidgetWin_1","processName":"chrome","activeDrag":null},{"type":"window","id":"462e2770-90ad-403e-86ab-3e668054fa2d","parentId":"a58a6ad0-1a0c-4196-81fd-b0d9dd97ed61","hasFocus":false,"tilingSize":0.5,"width":1271,"height":1378,"x":3843,"y":6,"state":{"type":"tiling"},"prevState":null,"displayState":"hidden","borderDelta":{"left":{"amount":0.0,"unit":"pixel"},"top":{"amount":0.0,"unit":"pixel"},"right":{"amount":0.0,"unit":"pixel"},"bottom":{"amount":0.0,"unit":"pixel"}},"floatingPlacement":{"left":3415,"top":75,"right":4265,"bottom":1315},"handle":133192,"title":".editorconfig - Visual Studio Code","className":"Chrome_WidgetWin_1","processName":"Code","activeDrag":null}],"childFocusOrder":["658617bd-b61f-46b2-869d-c12aed35c1f7","462e2770-90ad-403e-86ab-3e668054fa2d"],"hasFocus":false,"isDisplayed":false,"width":2548,"height":1378,"x":2566,"y":6,"tilingDirection":"horizontal"},{"type":"workspace","id":"be8350c7-d06e-4c1d-9347-4315844e1404","name":"9","displayName":"9","parentId":"7bab97e0-c9c4-4ac9-a479-bfb1c764a53f","children":[],"childFocusOrder":[],"hasFocus":false,"isDisplayed":true,"width":2548,"height":1378,"x":2566,"y":6,"tilingDirection":"horizontal"}]},"error":null,"success":true}
		// TODO: edge case after registering:
		// 2024/12/09 21:53:04 Received  {"messageType":"client_response","clientMessage":"sub -e focus_changed","data":{"subscriptionId":"08c01082-4b94-4731-b794-1b3147f180c1"},"error":null,"success":true}

		// TODO: Edge case after switching to a new workspace
		// 2024/12/09 21:54:13 Received  {"messageType":"event_subscription","data":{"eventType":"focus_changed","focusedContainer":{"type":"workspace","id":"9880001b-c245-4921-8748-f1774c532d4a","name":"7","displayName":"7","parentId":"f0bdd9e0-37b7-4572-b8ac-fa0203dd0886","children":[],"childFocusOrder":[],"hasFocus":true,"isDisplayed":true,"width":2548,"height":1378,"x":-2554,"y":6,"tilingDirection":"horizontal"}},"error":null,"subscriptionId":"08c01082-4b94-4731-b794-1b3147f180c1","success":true}
		workspaceName, ok := workspacesById[message.Data.FocusedContainer.ParentId]
		if !ok {
			workspaceName = workspacesById[message.Data.FocusedContainer.Id]
		}

		onWorkspaceSelected(workspaceName)
	}
}

var workspacesById = make(map[string]string)
var iconsByNumber = make(map[string][]byte)
var ready = false

func loadIcons() error {
	log.Println("Loading icons")
	for i := 0; i <= 9; i++ {
		// Icon source: https://github.com/urob/komotray
		icon, err := os.ReadFile(fmt.Sprint("assets/icons/", i, "-1.ico"))
		if err != nil {
			return err
		}

		iconsByNumber[strconv.Itoa(i)] = icon
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

func onReady() {
	ready = true
	systray.SetTooltip("Pretty awesome超级棒")
	systray.AddMenuItem("hello", "world")
	systray.SetIcon(iconsByNumber["0"])
}

func onExit() {
	// clean up here
}
