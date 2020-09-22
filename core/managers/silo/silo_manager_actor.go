package silo_manager

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

func (state *SiloManager) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *CreateProcessorMessage:

		newProcessorID, _ := state.CreateNewProcessor(msg)

		fmt.Printf("Created Processor\n")
		context.Respond(CreateProcessorOKMessage{PID: newProcessorID})

	case *CreateQueueMessage:
		newQueueID, _ := state.CreateNewQueue(msg)

		fmt.Printf("Created Queue\n")
		context.Respond(CreateQueueOKMessage{PID: newQueueID})

	case *StartSiloMessage:

		_ = state.StartSilo()
		context.Respond(StartSiloOKMessage{})

	case *StopSiloMessage:

		_ = state.StopSilo()
		context.Respond(StopSiloOKMessage{})

	case *AddInputQueueMessage:

		_ = state.AddInputQueue(msg.ProcessorID, msg.QueueID)
		context.Respond(AddInputQueueOKMessage{})

	case *AddOutputQueueMessage:

		_ = state.AddOutputQueue(msg.ProcessorID, msg.QueueID)
		context.Respond(AddInputQueueOKMessage{})

	case *GetProcessorsMessage:
		processors := state.GetProcessors()
		context.Respond(GetProcessorsOKMessage{
			Processors: processors,
		})

	}

}
