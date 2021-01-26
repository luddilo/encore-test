package hello

import (
	"context"
	"fmt"
)

type WorldResponse struct {
	// Message is the greeting response.
	Message string
}

// World responds with a familiar message.
//
// encore:api public
func World(ctx context.Context) (*WorldResponse, error) {
	msg := fmt.Sprintf("Hello, Simsim!")
	return &WorldResponse{Message: msg}, nil
}
