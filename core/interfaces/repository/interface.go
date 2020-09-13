package interface_repository

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfiles"
)

type RepositoryInterface interface {
	Write(flowfile *flowfiles.Flowfile, queueID uuid.UUID, value interface{}) error
	Read(flowfile *flowfiles.Flowfile) (interface{}, error)
	Commit(flowfile *flowfiles.Flowfile, savepoint bool) error
	Rollback(flowfile *flowfiles.Flowfile) error
	Remove(flowfile *flowfiles.Flowfile) error
}
