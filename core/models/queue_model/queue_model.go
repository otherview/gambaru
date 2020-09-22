package queue_model

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/google/uuid"
)

type Queue struct {
	ID      uuid.UUID
	ActorID *actor.PID
}
