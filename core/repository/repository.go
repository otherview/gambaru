package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfiles"
)

type MemoryRepository struct {
	flowfileData        map[uuid.UUID]interface{}
	lockFlowfileData    sync.RWMutex
	tmpFlowfileData     map[uuid.UUID]interface{}
	lockTmpFlowfileDate sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {

	return &MemoryRepository{flowfileData: map[uuid.UUID]interface{}{}, tmpFlowfileData: map[uuid.UUID]interface{}{}}
}

func (repo *MemoryRepository) Write(flowfile *flowfiles.Flowfile, value interface{}) error {

	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	repo.tmpFlowfileData[flowfile.ID] = value
	return nil
}

func (repo *MemoryRepository) Read(flowfile *flowfiles.Flowfile) (interface{}, error) {

	repo.lockFlowfileData.RLock()
	defer repo.lockFlowfileData.RUnlock()
	return repo.flowfileData[flowfile.ID], nil
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
