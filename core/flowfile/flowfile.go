package flowfiles

import "github.com/google/uuid"

type Flowfile struct {
	ID uuid.UUID
}

func NewFlowfile() *Flowfile {

	return &Flowfile{
		ID: uuid.New(),
	}
}
