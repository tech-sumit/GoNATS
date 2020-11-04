package models

type CountCharsRequest struct {
	String string `json:"string"`
}

type CountCharsResponse struct {
	Status bool `json:"status"`
	Count  int  `json:"count,omitempty"`
}
