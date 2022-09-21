package io

// MessageOutput is a response with a simple message
type MessageOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// HTTPError is returned for errors within HTTP handlers
type HTTPError struct {
	Error string `json:"error"`
}
