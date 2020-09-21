package flow_model

import (
	"github.com/google/uuid"
)

type Queue struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Type string    `json:"type,omitempty"`
}

type Queues []Queue

type Config struct {
	ExecutionTime string `json:"executionTime,omitempty"`
	SavePoint     bool   `json:"savePoint,omitempty"`
}

type Processor struct {
	ID           uuid.UUID   `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	InputQueues  []uuid.UUID `json:"inputQueues,omitempty"`
	OutputQueues []uuid.UUID `json:"outputQueues,omitempty"`
	Config       *Config     `json:"config,omitempty"`
}

type Processors []Processor

type Structure struct {
	Queues     `json:"queues"`
	Processors `json:"processors"`
}
