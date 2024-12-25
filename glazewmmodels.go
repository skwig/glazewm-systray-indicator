package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// https://github.com/glzr-io/glazewm-js/blob/main/src/types/events

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

type Window struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	ParentId string `json:"parentId"`
	HasFocus bool   `json:"hasFocus"`
	// Note that there are other properties, but they are not relevant for our usecase
}

// TODO: Base events
type FocusChangedEvent struct {
	EventType        string                  `json:"eventType"`
	FocusedContainer FocusedContainerWrapper `json:"focusedContainer"`
}

type FocusedContainer interface {
	GetFocusedContainerType() string
}

type FocusedContainerWrapper struct {
	Value FocusedContainer
}

func (window Window) GetFocusedContainerType() string {
	return window.Type
}

func (workspace Workspace) GetFocusedContainerType() string {
	return workspace.Type
}

func (wrapper *FocusedContainerWrapper) UnmarshalJSON(data []byte) error {
	var distriminator struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(data, &distriminator); err != nil {
		return err
	}

	switch distriminator.Type {
	case "workspace":
		var workspace Workspace
		if err := json.Unmarshal(data, &workspace); err != nil {
			return err
		}
		wrapper.Value = workspace
	case "window":
		var window Window
		if err := json.Unmarshal(data, &window); err != nil {
			return err
		}
		wrapper.Value = window
	default:
		return errors.New(fmt.Sprintf("unknown type: %s", distriminator.Type))
	}

	return nil
}
