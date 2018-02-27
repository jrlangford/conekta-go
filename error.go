package conekta

import (
	"encoding/json"
)

// ErrorDetails gives especific information about an error
type ErrorDetails struct {
	DebugMessage string `json:"debug_message"`
	Message      string `json:"message"`
	Param        string `json:"param"`
	Code         string `json:"code"`
}

// Error describes Conekta errors
type Error struct {
	ErrorType string         `json:"type"`
	LogID     string         `json:"log_id"`
	Details   []ErrorDetails `json:"details"`
}

func getConectionError() *Error {
	return &Error{
		ErrorType: "error.requestor.connection_purchaser",
		Details: []ErrorDetails{
			{
				Message: "error.requestor.connection",
			},
		},
	}
}

func getAPIError(b []byte) *Error {
	e := Error{}
	json.Unmarshal(b, &e)
	return &e
}
