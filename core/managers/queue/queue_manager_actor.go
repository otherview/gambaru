package queue_manager

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func (state *QueueManager) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *ReadQueueItemMessage:
		queueItem := state.ReadQueueItem()
		context.Respond(ReadQueueItemOKMessage{QueueItem: queueItem})
	case *WriteQueueItemMessage:
		_ = state.WriteQueueItem(msg.QueueItem)
		context.Respond(WriteQueueItemOKMessage{})
	}

}
