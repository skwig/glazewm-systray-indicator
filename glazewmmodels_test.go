package main

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFocusChangedWorkspace(t *testing.T) {
	jsonData := `
    {
      "messageType": "event_subscription",
      "data": {
        "eventType": "focus_changed",
        "focusedContainer": {
          "type": "workspace",
          "id": "7ff3ae2b-d00d-48b5-b216-9903ae14373c",
          "name": "7",
          "displayName": "7",
          "parentId": "c3cdcf92-f07b-41b7-8273-f388c6a0e40b",
          "children": [],
          "childFocusOrder": [],
          "hasFocus": true,
          "isDisplayed": true,
          "width": 2860,
          "height": 1684,
          "x": 10,
          "y": 10,
          "tilingDirection": "horizontal"
        }
      },
      "error": null,
      "subscriptionId": "efbc0137-19cf-40ba-8b8d-c7b51c3b6726",
      "success": true
    }
    `

	var message GlazeWmMessage[FocusChangedEvent]
	if err := json.Unmarshal([]byte(jsonData), &message); err != nil {
		t.Fatal("Error:", err)
		return
	}

	diff := cmp.Diff(message.Data.FocusedContainer.Value,
		Workspace{
			Type:        "workspace",
			Id:          "7ff3ae2b-d00d-48b5-b216-9903ae14373c",
			Name:        "7",
			DisplayName: "7",
			ParentId:    "c3cdcf92-f07b-41b7-8273-f388c6a0e40b",
			HasFocus:    true,
			IsDisplayed: true})

	if diff != "" {
		t.Fatal(diff)
	}
}

func TestFocusChangedWindow(t *testing.T) {
	jsonData := `
    {
      "messageType": "event_subscription",
      "data": {
        "eventType": "focus_changed",
        "focusedContainer": {
          "type": "window",
          "id": "db75115c-f333-46fb-a87b-eef291a530c1",
          "parentId": "8f0ea4a4-17c0-4db6-aea4-f5b22a5c0726",
          "hasFocus": true,
          "tilingSize": 1.0,
          "width": 2860,
          "height": 1684,
          "x": 10,
          "y": 10,
          "state": {
            "type": "tiling"
          },
          "prevState": {
            "type": "minimized"
          },
          "displayState": "showing",
          "borderDelta": {
            "left": {
              "amount": 0.0,
              "unit": "pixel"
            },
            "top": {
              "amount": 0.0,
              "unit": "pixel"
            },
            "right": {
              "amount": 0.0,
              "unit": "pixel"
            },
            "bottom": {
              "amount": 0.0,
              "unit": "pixel"
            }
          },
          "floatingPlacement": {
            "left": 725,
            "top": 87,
            "right": 2156,
            "bottom": 1617
          },
          "handle": 985126,
          "title": "Comparing and Diffing Objects in Go Tests – Tabstop - Google Chrome",
          "className": "Chrome_WidgetWin_1",
          "processName": "chrome",
          "activeDrag": null
        }
      },
      "error": null,
      "subscriptionId": "efbc0137-19cf-40ba-8b8d-c7b51c3b6726",
      "success": true
    }
    `

	var message GlazeWmMessage[FocusChangedEvent]
	if err := json.Unmarshal([]byte(jsonData), &message); err != nil {
		t.Fatal("Error:", err)
		return
	}

	diff := cmp.Diff(message.Data.FocusedContainer.Value,
		Window{
			Type:     "window",
			Id:       "db75115c-f333-46fb-a87b-eef291a530c1",
			ParentId: "8f0ea4a4-17c0-4db6-aea4-f5b22a5c0726",
			HasFocus: true})

	if diff != "" {
		t.Fatal(diff)
	}
}

func TestWorkspaceActivated(t *testing.T) {
	jsonData := `
    {
      "messageType": "event_subscription",
      "data": {
        "eventType": "workspace_activated",
        "activatedWorkspace": {
          "type": "workspace",
          "id": "f6279441-b65b-4f80-9120-7adede0e4b89",
          "name": "8",
          "displayName": "8",
          "parentId": "9c16720d-2ba9-4f11-ab79-913b1f2afe1d",
          "children": [],
          "childFocusOrder": [],
          "hasFocus": false,
          "isDisplayed": false,
          "width": 2860,
          "height": 1684,
          "x": 10,
          "y": 10,
          "tilingDirection": "horizontal"
        }
      },
      "error": null,
      "subscriptionId": "efbc0137-19cf-40ba-8b8d-c7b51c3b6726",
      "success": true
    }
    `

	var message GlazeWmMessage[WorkspaceActivatedEvent]
	if err := json.Unmarshal([]byte(jsonData), &message); err != nil {
		t.Fatal("Error:", err)
		return
	}

	diff := cmp.Diff(message.Data.ActivatedWorkspace,
		Workspace{
			Type:        "workspace",
			Id:          "f6279441-b65b-4f80-9120-7adede0e4b89",
			Name:        "8",
			DisplayName: "8",
			ParentId:    "9c16720d-2ba9-4f11-ab79-913b1f2afe1d",
			HasFocus:    false,
			IsDisplayed: false})

	if diff != "" {
		t.Fatal(diff)
	}
}
