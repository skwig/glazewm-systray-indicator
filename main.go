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
	connection.WriteMessage(1, []byte("sub -e focus_changed"))

	// Kickstart the state
	connection.WriteMessage(1, []byte("query monitors"))

	for {
		var message MessageWrapper
		_, buf, err := connection.ReadMessage()
		if err != nil {
			log.Println("Read error", err)
			continue
		}

		err = json.Unmarshal(buf, &message)
		if err != nil {
			log.Println("Unmarshal error", err)
			continue
		}

		switch t := message.Value.(type) {
		case EventMessage:
			connection.WriteMessage(1, []byte("query monitors"))
		case ResponseMessage:
			for _, monitor := range t.Data.Monitors {
				if !monitor.HasFocus {
					continue
				}

				for _, workspace := range monitor.Children {
					if !workspace.HasFocus {
						continue
					}

					onWorkspaceSelected(workspace.Name)
				}
			}
		default:
			log.Println("Received ", string(buf))
		}
	}
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
