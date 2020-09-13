package repository

import (
	"encoding/json"

	"github.com/google/uuid"
)

type repositoryFlowfile struct {
	Data       interface{}
	QueueID    uuid.UUID
	FlowfileID uuid.UUID
	committed  bool
}

func newRepositoryFlowfile(flowfileID uuid.UUID, queueID uuid.UUID, data interface{}) *repositoryFlowfile {

	return &repositoryFlowfile{
		Data:       data,
		QueueID:    queueID,
		FlowfileID: flowfileID,
		committed:  false,
	}
}

func (f *repositoryFlowfile) Bytes() []byte {

	data, err := json.Marshal(&f)
	if err != nil {
		// TODO yep
		panic(err)
	}

	return data
}

func (f *repositoryFlowfile) HasBeenWritten() bool {
	return f.committed
}

func (f *repositoryFlowfile) Remove() {

}
