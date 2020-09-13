package flow_model

import (
	"github.com/google/uuid"
)

type Queues []struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Type string    `json:"type,omitempty"`
}

type Config struct {
	ExecutionTime string `json:"executionTime,omitempty"`
	SavePoint     bool   `json:"savePoint,omitempty"`
}

type Processors []struct {
	ID           uuid.UUID   `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	InputQueues  []uuid.UUID `json:"inputQueues,omitempty"`
	OutputQueues []uuid.UUID `json:"outputQueues,omitempty"`
	Config       *Config     `json:"config,omitempty"`
}

type Structure struct {
	Queues     `json:"queues"`
	Processors `json:"processors"`
}
