// Package io contains inputs and outputs
package io

// AllocateInput is the input for the /allocate-memory endpoint
type AllocateInput struct {
	ArrayLength     *string `json:"arrayLength"`
	BytesPerElement *string `json:"bytesPerElement"`
}
