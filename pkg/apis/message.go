package apis

import (
	"encoding/json"
	"net/http"
)

// Message is the object return from every request
type Message struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Metadata string `json:"metadata"`
}

// JSON return the JSON representation of the Message
func (m *Message) JSON() string {
	marshaled, _ := json.Marshal(m)

	res := string(marshaled)
	return res
}

// NewError creates a new error Message
func NewError(code int, error error) *Message {
	return &Message{
		Status:   "error",
		Code:     code,
		Metadata: error.Error(),
	}
}

// NewSuccess creates a new success Message
func NewSuccess(metadata string) *Message {
	return &Message{
		Status:   "success",
		Code:     http.StatusOK,
		Metadata: metadata,
	}
}
