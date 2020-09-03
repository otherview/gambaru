package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfile"
)

type Repository struct {
	flowfileData map[uuid.UUID]interface{}
	lock         sync.RWMutex
}

func NewRepository() *Repository {

	return &Repository{flowfileData: map[uuid.UUID]interface{}{}}
}

func (repo *Repository) Write(flowfile *flowfiles.Flowfile, value interface{}) error {

	repo.lock.Lock()
	defer repo.lock.Unlock()
	repo.flowfileData[flowfile.ID] = value
	return nil
}

func (repo *Repository) Read(flowfile *flowfiles.Flowfile) (interface{}, error) {

	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.flowfileData[flowfile.ID], nil
}
