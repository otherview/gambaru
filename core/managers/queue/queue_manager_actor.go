package queue_manager

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	queue_manager_messages "github.com/otherview/gambaru/core/managers/queue/messages"
)

func (state *QueueManager) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *queue_manager_messages.ReadQueueItemMessage:
		queueItem := state.ReadQueueItem()
		context.Respond(queue_manager_messages.ReadQueueItemOKMessage{QueueItem: queueItem})
	case *queue_manager_messages.WriteQueueItemMessage:
		_ = state.WriteQueueItem(msg.QueueItem)
		context.Respond(queue_manager_messages.WriteQueueItemOKMessage{})
	}

}
