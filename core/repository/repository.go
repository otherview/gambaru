package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfiles"
)

type MemoryRepository struct {
	flowfileData        map[uuid.UUID]*repositoryFlowfile
	lockFlowfileData    sync.RWMutex
	tmpFlowfileData     map[uuid.UUID]*repositoryFlowfile
	lockTmpFlowfileDate sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {

	return &MemoryRepository{flowfileData: map[uuid.UUID]*repositoryFlowfile{}, tmpFlowfileData: map[uuid.UUID]*repositoryFlowfile{}}
}

func (repo *MemoryRepository) Write(flowfile *flowfiles.Flowfile, queueID uuid.UUID, value interface{}) error {

	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	repo.tmpFlowfileData[flowfile.ID] = newRepositoryFlowfile(flowfile.ID, queueID, value)
	return nil
}

func (repo *MemoryRepository) Read(flowfile *flowfiles.Flowfile) (interface{}, error) {

	repo.lockFlowfileData.RLock()
	defer repo.lockFlowfileData.RUnlock()
	return repo.flowfileData[flowfile.ID].Data, nil
}

func (repo *MemoryRepository) Commit(flowfile *flowfiles.Flowfile) error {

	repo.lockFlowfileData.Lock()
	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	defer repo.lockFlowfileData.Unlock()
	repo.flowfileData[flowfile.ID] = repo.tmpFlowfileData[flowfile.ID]
	return nil
}

func (repo *MemoryRepository) Rollback(flowfile *flowfiles.Flowfile) error {

	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	delete(repo.flowfileData, flowfile.ID)
	return nil
}

func (repo *MemoryRepository) Remove(flowfile *flowfiles.Flowfile) error {
	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	delete(repo.flowfileData, flowfile.ID)

	return nil
}
