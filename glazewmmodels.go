package main

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
