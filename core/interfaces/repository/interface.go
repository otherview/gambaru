package interface_repository

import "github.com/otherview/gambaru/core/flowfiles"

type RepositoryInterface interface {
	Write(flowfile *flowfiles.Flowfile, value interface{}) error
	Read(flowfile *flowfiles.Flowfile) (interface{}, error)
	Commit(flowfile *flowfiles.Flowfile) error
	Rollback(flowfile *flowfiles.Flowfile) error
}
