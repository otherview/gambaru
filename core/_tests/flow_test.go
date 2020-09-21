package _tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowchart"
	"github.com/otherview/gambaru/core/models/flow_model"
	"github.com/otherview/gambaru/core/silo"
)

func TestGeneratorAndLoggerProcessors(t *testing.T) {

	queue := flow_model.Queue{
		ID:   uuid.New(),
		Type: "SimpleQueue",
	}

	generatorProcessor := flow_model.Processor{
		ID:           uuid.New(),
		Type:         "SimpleTextGeneratorProcessor",
		InputQueues:  nil,
		OutputQueues: []uuid.UUID{queue.ID},
		Config:       nil,
	}

	logProcessor := flow_model.Processor{
		ID:           uuid.New(),
		Type:         "SimpleLogProcessor",
		InputQueues:  []uuid.UUID{queue.ID},
		OutputQueues: nil,
		Config:       nil,
	}

	siloInstance := silo.NewSilo()
	flow := flowchart.NewFlow(siloInstance)

	err := flow.CreateProcessor(&generatorProcessor)
	if err != nil {
		t.Errorf("error creating a procerror, %v", generatorProcessor)
	}

	err = flow.CreateProcessor(&logProcessor)
	if err != nil {
		t.Errorf("error creating a procerror, %v", generatorProcessor)
	}

	err = flow.CreateQueue(&queue)
	if err != nil {
		t.Errorf("error creating a queue, %v", queue)
	}
}
