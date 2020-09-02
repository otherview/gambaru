package flow

import (
	"github.com/google/uuid"
)

type Queues []struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Type string    `json:"type,omitempty"`
}

type Processors []struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	InputQueue  uuid.UUID `json:"inputQueue,omitempty"`
	OutputQueue uuid.UUID `json:"outputQueue,omitempty"`
}

type Structure struct {
	Queues     `json:"queues"`
	Processors `json:"processors"`
}
