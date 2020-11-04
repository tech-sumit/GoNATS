package models

type ReverseStringRequest struct {
	String string `json:"string"`
}

type ReverseStringResponse struct {
	Status bool   `json:"status"`
	String string `json:"string,omitempty"`
}
