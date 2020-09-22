package processor_manager

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func (state *ProcessorManager) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *StartProcessorMessage:
		_ = state.StartProcessor()
		context.Respond(StartProcessorOKMessage{})
	case *StopProcessorMessage:
		_ = state.StopProcessor()
		context.Respond(StopProcessorOKMessage{})

	case *AddInputQueue:
		_ = state.AddInputQueue(msg)
		context.Respond(AddInputOKQueue{})

	case *AddOutputQueue:
		_ = state.AddOutputQueue(msg)
		context.Respond(AddOutputOKQueue{})

	case *GetProcessorInfoMessage:
		procID, procType, inputQueues, outputQueues := state.GetProcessorInfo()
		context.Respond(GetProcessorInfoOKMessage{
			ID:           procID,
			Type:         procType,
			InputQueues:  inputQueues,
			OutputQueues: outputQueues,
		})
	}

}
