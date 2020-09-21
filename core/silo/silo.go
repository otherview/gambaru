package silo

import (
	"time"

	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"
	interface_queue "github.com/otherview/gambaru/core/interfaces/queue"

	"github.com/otherview/gambaru/core/repository"

	"github.com/otherview/gambaru/lib/queues"

	silo_manager "github.com/otherview/gambaru/core/managers/silo"

	"github.com/AsynkronIT/protoactor-go/actor"
	procs "github.com/otherview/gambaru/lib/processors"
)

// Silo is the running instance of gambaru
type Silo struct {
	siloPID              *actor.PID
	context              *actor.RootContext
	repository           *repository.WALRepository
	registeredProcessors map[string]interface_processor.ProcessorInterface
	registeredQueues     map[string]interface_queue.QueueInterface
}

// NewSilo creates a new silo
func NewSilo() *Silo {

	// TODO this should be a singleton
	// Creates the Silo Manager actor
	props := actor.PropsFromProducer(func() actor.Actor { return silo_manager.NewSiloManager() })
	rootContext := actor.EmptyRootContext
	pid := rootContext.Spawn(props)

	availableProcessors := map[string]interface_processor.ProcessorInterface{}
	availableProcessors["SimpleLogProcessor"] = procs.NewSimpleLogProcessor()
	availableProcessors["SimpleTextGeneratorProcessor"] = procs.NewSimpleTextGeneratorProcessor()
	availableProcessors["SimpleTextFinderProcessor"] = procs.NewSimpleTextFinderProcessor()

	availableQueues := map[string]interface_queue.QueueInterface{}
	availableQueues["SimpleQueue"] = queues.NewBufferQueue()

	return &Silo{
		siloPID:              pid,
		context:              rootContext,
		registeredProcessors: availableProcessors,
		registeredQueues:     availableQueues,
	}
}

// Start commences all the configured processors
func (silo *Silo) Start() error {

	_, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.StartSiloMessage{}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}

	return nil
}

// Stop stops all the configured processors
func (silo *Silo) Stop() {

	_, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.StopSiloMessage{}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}
}

// GetRegisteredProcessor returns a registered type of processor
func (silo *Silo) GetRegisteredProcessor(processorTypeName string) interface_processor.ProcessorInterface {
	return silo.registeredProcessors[processorTypeName]
}

// GetRegisteredQueue returns a registered type of queue
func (silo *Silo) GetRegisteredQueue(queueTypeName string) interface_queue.QueueInterface {

	return silo.registeredQueues[queueTypeName]
}
