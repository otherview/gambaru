package silo

import (
	"time"

	"github.com/otherview/gambaru/core/repository"

	"github.com/otherview/gambaru/core"

	"github.com/otherview/gambaru/lib/queues"

	silo_manager "github.com/otherview/gambaru/core/managers/silo"

	"github.com/AsynkronIT/protoactor-go/actor"
	procs "github.com/otherview/gambaru/lib/processors"
)

// Silo is the running instance of gambaru
type Silo struct {
	siloPID              *actor.PID
	context              *actor.RootContext
	repository           *repository.Repository
	registeredProcessors map[string]core.ProcessorInterface
	registeredQueues     map[string]core.QueueInterface
}

// NewSilo creates a new silo
func NewSilo() *Silo {

	// TODO this should be a singleton
	// Creates the Silo Manager actor
	props := actor.PropsFromProducer(func() actor.Actor { return silo_manager.NewSiloManager() })
	rootContext := actor.EmptyRootContext
	pid := rootContext.Spawn(props)

	newRepository := repository.NewRepository()

	availableProcessors := map[string]core.ProcessorInterface{}
	availableProcessors["SimpleLogProcessor"] = procs.NewSimpleLogProcessor(newRepository)
	availableProcessors["SimpleTextGeneratorProcessor"] = procs.NewSimpleTextGeneratorProcessor(newRepository)

	availableQueues := map[string]core.QueueInterface{}
	availableQueues["SimpleQueue"] = queues.NewSimpleQueue()

	return &Silo{
		siloPID:              pid,
		context:              rootContext,
		registeredProcessors: availableProcessors,
		registeredQueues:     availableQueues,
		repository:           newRepository,
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
func (silo *Silo) GetRegisteredProcessor(processorTypeName string) core.ProcessorInterface {
	return silo.registeredProcessors[processorTypeName]
}

// GetRegisteredQueue returns a registered type of queue
func (silo *Silo) GetRegisteredQueue(queueTypeName string) core.QueueInterface {

	return silo.registeredQueues[queueTypeName]
}
