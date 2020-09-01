package flow

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/silo"
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

type FlowModel struct {
	Queues     `json:"queues"`
	Processors `json:"processors"`
}

func (m *FlowModel) CreateFlow(silo *silo.Silo, flow *FlowModel) {

	for _, queue := range flow.Queues {
		silo.CreateQueue(nil, queue.ID)
	}

	for _, processor := range flow.Processors {
		// TODO check if it exists :o
		procType := silo.GetProcessorType(processor.Type)
		silo.CreateProcessor(procType, processor.ID)
		silo.AddInputQueue(processor.ID, processor.InputQueue)
		silo.AddOutputQueue(processor.ID, processor.OutputQueue)
	}

}
