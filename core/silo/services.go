package silo

import (
	"time"

	silo_manager "github.com/otherview/gambaru/core/managers/silo"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/processors"
	"github.com/otherview/gambaru/core/queues"
	procs "github.com/otherview/gambaru/lib/processors"
)

type Silo struct {
	siloPID             *actor.PID
	context             *actor.RootContext
	availableProcessors map[string]processors.ProcessorInterface
}

func NewSilo() *Silo {

	// TODO this should be a singleton
	// Creates the Silo Manager actor
	props := actor.PropsFromProducer(func() actor.Actor { return silo_manager.NewSiloManager() })
	rootContext := actor.EmptyRootContext
	pid := rootContext.Spawn(props)

	availableProcessors := map[string]processors.ProcessorInterface{}
	availableProcessors["SimpleLogProcessor"] = procs.NewSimpleLogProcessor()
	availableProcessors["SimpleTextGeneratorProcessor"] = procs.NewSimpleTextGeneratorProcessor()

	return &Silo{
		siloPID:             pid,
		context:             rootContext,
		availableProcessors: availableProcessors,
	}
}

func (silo *Silo) CreateProcessor(processor processors.ProcessorInterface, processorID uuid.UUID) uuid.UUID {

	createdProcessorMessage, err := silo.context.RequestFuture(silo.siloPID,
		&silo_manager.CreateProcessorMessage{
			Processor: &processor,
			ID:        processorID,
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

func (silo *Silo) CreateQueue(queue *queues.QueueInterface, queueID uuid.UUID) uuid.UUID {

	createdQueueMessage, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.CreateQueueMessage{
		Queue: queue,
		ID:    queueID,
	}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}

	createdQueue, ok := createdQueueMessage.(silo_manager.CreateQueueOKMessage)
	if !ok {
		return uuid.Nil
	}

	return createdQueue.PID
}

func (silo *Silo) Start() error {

	_, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.StartSiloMessage{}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}

	return nil
}

func (silo *Silo) Stop() {

	_, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.StopSiloMessage{}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}
}

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

func (silo *Silo) GetProcessorType(processorTypeName string) processors.ProcessorInterface {

	return silo.availableProcessors[processorTypeName]
}
