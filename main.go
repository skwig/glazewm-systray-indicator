package main

import (
	"log"
	"os"

	"github.com/getlantern/systray"
	// "github.com/gorilla/websocket"
	// "github.com/gorilla/websocket"
)

// https://github.com/getlantern/systray
// https://github.com/gorilla/websocket
// https://github.com/joshprk/komotray
// https://github.com/glzr-io/glazewm-js/tree/main/src/types
// https://discord.com/channels/1041662798196908052/1041662798813466707/1295247854276972565

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

func main() {
	systray.Run(onReady, onExit)

	// // TODO: How to have err not freak out here?
	// c, _, err0 := websocket.DefaultDialer.Dial("ws://localhost:6123", nil)
	// if err0 != nil {
	// 	log.Println("read:", err0)
	// 	return
	// }
	//
	// c.WriteMessage(1, []byte("query workspaces"))
	// var queryResponse GlazeWmMessage[Workspaces]
	// err := c.ReadJSON(&queryResponse)
	// if err != nil {
	// 	log.Println("read:", err)
	// 	return
	// }
	//
	// workspacesById := make(map[string]string)
	// for _, w := range queryResponse.Data.Workspaces {
	// 	if w.HasFocus {
	// 		onWorkspaceSelected(w.Name)
	// 	}
	// 	workspacesById[w.Id] = w.Name
	// }
	//
	// c.WriteMessage(1, []byte("sub -e focus_changed"))
	// for {
	// 	var message GlazeWmMessage[FocusChangedEvent]
	// 	err := c.ReadJSON(&message)
	// 	if err != nil {
	// 		log.Println("read:", err)
	// 		continue
	// 	}
	//
	// 	workspaceName, ok := workspacesById[message.Data.FocusedContainer.ParentId]
	// 	if !ok {
	// 		workspaceName = workspacesById[message.Data.FocusedContainer.Id]
	// 	}
	//
	// 	onWorkspaceSelected(workspaceName)
	// }
}

func onWorkspaceSelected(workspaceName string) {
	log.Println(workspaceName)
	// systray.SetTooltip(workspaceName)
}

func onReady() {
	icon, _ := os.ReadFile("assets/icons/1-1.ico")
	// systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	systray.AddMenuItem("hello", "world")
	systray.SetIcon(icon)
	// mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	//
	// Sets the icon of a menu item. Only available on Mac and Windows.
	// mQuit.SetIcon(icon.Data)
}

func onExit() {
	// clean up here
}
