package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfiles"
	"github.com/tidwall/wal"
)

type WALRepository struct {
	flowfileData        map[uuid.UUID]*repositoryFlowfile
	lockFlowfileData    sync.RWMutex
	tmpFlowfileData     map[uuid.UUID]*repositoryFlowfile
	lockTmpFlowfileDate sync.RWMutex
	lockWriterPosition  sync.RWMutex
	writerPosition      uint64
	writerPositionMap   map[uuid.UUID]uint64
	logWriter           *wal.Log
}

func NewWALRepository() *WALRepository {

	// open a new log file
	logWriter, err := wal.Open("writelogs", nil)
	if err != nil {
		// TODO Yep
		panic(err)
	}

	return &WALRepository{
		flowfileData:      map[uuid.UUID]*repositoryFlowfile{},
		tmpFlowfileData:   map[uuid.UUID]*repositoryFlowfile{},
		writerPosition:    uint64(0),
		writerPositionMap: map[uuid.UUID]uint64{},
		logWriter:         logWriter,
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
		return repo.flowfileData[flowfile.ID].Data, nil
	}
	return nil, nil
}

func (repo *WALRepository) Commit(flowfile *flowfiles.Flowfile) error {

	var err error

	repo.lockFlowfileData.Lock()
	repo.lockTmpFlowfileDate.Lock()
	defer repo.lockTmpFlowfileDate.Unlock()
	defer repo.lockFlowfileData.Unlock()
	repo.flowfileData[flowfile.ID] = repo.tmpFlowfileData[flowfile.ID]

	if !repo.flowfileData[flowfile.ID].HasBeenWritten() {

		err = repo.logWriter.Write(repo.newPosition(flowfile.ID), repo.flowfileData[flowfile.ID].Bytes())
		if err != nil {
			// TODO Yep
			panic(err)
		}
		repo.flowfileData[flowfile.ID].Commit()
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

	if pos, ok := repo.writerPositionMap[flowfile.ID]; ok {
		err := repo.logWriter.TruncateBack(pos)
		if err != nil {
			// TODO Yep
			panic(err)
		}
		err = repo.logWriter.TruncateFront(pos)
		if err != nil {
			// TODO Yep
			panic(err)
		}
	}

	return nil
}

func (repo *WALRepository) newPosition(flowfileID uuid.UUID) uint64 {
	repo.lockWriterPosition.Lock()
	defer repo.lockWriterPosition.Unlock()
	repo.writerPosition += 1
	repo.writerPositionMap[flowfileID] = repo.writerPosition

	return repo.writerPosition
}
