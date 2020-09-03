package repository

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfile"
)

type Repository struct {
	flowfileData map[uuid.UUID]interface{}
}

func NewRepository() *Repository {

	return &Repository{flowfileData: map[uuid.UUID]interface{}{}}
}

func (repo *Repository) Write(flowfile *flowfiles.Flowfile, value interface{}) error {

	repo.flowfileData[flowfile.ID] = value
	return nil
}

func (repo *Repository) Read(flowfile *flowfiles.Flowfile) (interface{}, error) {

	return repo.flowfileData[flowfile.ID], nil
}
