package silo

import (
	"time"

	"github.com/otherview/gambaru/core/models/flow_model"

	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"

	"github.com/google/uuid"
	silo_manager "github.com/otherview/gambaru/core/managers/silo"
)

// CreateProcessor returns the id of a created Processor
func (silo *Silo) CreateProcessor(processor interface_processor.ProcessorInterface, processorID uuid.UUID, config *flow_model.Config) uuid.UUID {

	// TODO theres a better way to do this
	processorConfig := &flow_model.Config{
		ExecutionTime: "1s",
		SavePoint:     false,
	}
	if config != nil {
		if config.ExecutionTime != "" {
			processorConfig.ExecutionTime = config.ExecutionTime
		}
		processorConfig.SavePoint = config.SavePoint
	}

	createdProcessorMessage, err := silo.context.RequestFuture(silo.siloPID,
		&silo_manager.CreateProcessorMessage{
			Processor: processor,
			ID:        processorID,
			Config:    processorConfig,
		}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}

	createdProcessor, ok := createdProcessorMessage.(silo_manager.CreateProcessorOKMessage)
	if !ok {
		return uuid.Nil
	}

	return createdProcessor.PID
}

// AddOutputQueue attaches an Output Queue to a processor
func (silo *Silo) AddOutputQueue(processorID uuid.UUID, queueID uuid.UUID) error {

	addOutputQueueMessage, err := silo.context.RequestFuture(silo.siloPID,
		&silo_manager.AddOutputQueueMessage{
			ProcessorID: processorID,
			QueueID:     queueID,
		},
		5*time.Second).Result()

	if err != nil {
		// TODO: yep
		panic(err)
	}

	_, ok := addOutputQueueMessage.(silo_manager.AddOutputQueueMessage)
	if !ok {
		// TODO make an error definition here between silo and service
		return err
	}

	return nil
}

// AddInputQueue attaches an Input Queue to a processor
func (silo *Silo) AddInputQueue(processorID uuid.UUID, queueID uuid.UUID) error {

	addInputQueueMessage, err := silo.context.RequestFuture(silo.siloPID,
		&silo_manager.AddInputQueueMessage{
			ProcessorID: processorID,
			QueueID:     queueID,
		},
		5*time.Second).Result()

	if err != nil {
		// TODO: yep
		panic(err)
	}

	_, ok := addInputQueueMessage.(silo_manager.AddInputQueueMessage)
	if !ok {
		// TODO make an error definition here between silo and service
		return err
	}

	return nil
}
