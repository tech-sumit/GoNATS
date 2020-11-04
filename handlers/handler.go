package handlers

import (
	"nats-template/core/models"
	"nats-template/lib"
)

type ClientHandler struct{}

func (clientHandler ClientHandler) countChars(request models.CountCharsRequest) models.CountCharsResponse {
	if request.String == "" {
		return models.CountCharsResponse{
			Status: false,
			Count:  0,
		}
	}
	return models.CountCharsResponse{
		Status: false,
		Count:  len(request.String),
	}
}

func (clientHandler ClientHandler) reverseString(request models.ReverseStringRequest) models.ReverseStringResponse {
	if request.String == "" {
		return models.ReverseStringResponse{Status: false}
	} else {
		return models.ReverseStringResponse{
			Status: true,
			String: lib.ReverseString(request.String),
		}
	}
}
