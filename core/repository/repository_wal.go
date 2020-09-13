package repository

import (
	"fmt"
	"sync"

	"github.com/otherview/gambaru/core/repository/provider"
	"github.com/tidwall/wal"

	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfiles"
)

type WALRepository struct {
	flowfileData        map[uuid.UUID]*repositoryFlowfile
	lockFlowfileData    sync.RWMutex
	tmpFlowfileData     map[uuid.UUID]*repositoryFlowfile
	lockTmpFlowfileDate sync.RWMutex
	lockWriterPosition  sync.RWMutex
	repoProvider        *provider.MemoryRepoProvider
	writerPosition      uint64
	writerPositionMap   map[uuid.UUID]uint64
	logWriter           *wal.Log
}

func NewWALRepository() *WALRepository {

	return &WALRepository{
		flowfileData:    map[uuid.UUID]*repositoryFlowfile{},
		tmpFlowfileData: map[uuid.UUID]*repositoryFlowfile{},
		repoProvider:    provider.NewMemoryRepoProvider(),
	}
}

func (repo *WALRepository) Write(flowfile *flowfiles.Flowfile, queueID uuid.UUID, value interface{}) error {

	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	repo.tmpFlowfileData[flowfile.ID] = newRepositoryFlowfile(flowfile.ID, queueID, value)
	return nil
}

func (repo *WALRepository) Read(flowfile *flowfiles.Flowfile) (interface{}, error) {

	repo.lockFlowfileData.RLock()
	defer repo.lockFlowfileData.RUnlock()
	if _, ok := repo.flowfileData[flowfile.ID]; ok {
		if repo.flowfileData[flowfile.ID] == nil {
			panic("derp")
		}
		return repo.flowfileData[flowfile.ID].Data, nil
	}
	return nil, nil
}

func (repo *WALRepository) Commit(flowfile *flowfiles.Flowfile, savepoint bool) error {

	var err error

	repo.lockFlowfileData.Lock()
	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	defer repo.lockFlowfileData.Unlock()
	repo.flowfileData[flowfile.ID] = repo.tmpFlowfileData[flowfile.ID]

	if savepoint {
		err := repo.repoProvider.Write(flowfile.ID, repo.flowfileData[flowfile.ID].Bytes())
		if err != nil {
			// TODO Yep
			fmt.Println("derp")
			panic(err)
		}
	}

	return err
}

func (repo *WALRepository) Rollback(flowfile *flowfiles.Flowfile) error {

	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	delete(repo.flowfileData, flowfile.ID)
	return nil
}

func (repo *WALRepository) Remove(flowfile *flowfiles.Flowfile) error {
	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	delete(repo.flowfileData, flowfile.ID)

	err := repo.repoProvider.Remove(flowfile.ID)
	if err != nil {
		// TODO Yep
		fmt.Println("derp")
		panic(err)
	}
	return nil
}
