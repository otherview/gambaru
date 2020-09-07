package sessions

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/otherview/gambaru/core/flowfiles"
)

type operation struct {
	flowfile  *flowfiles.Flowfile
	destQueue *actor.PID
}

func newOperation(flowfile *flowfiles.Flowfile, destQueue *actor.PID) *operation {

	return &operation{flowfile: flowfile, destQueue: destQueue}
}
