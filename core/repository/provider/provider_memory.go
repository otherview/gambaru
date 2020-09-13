package provider

import (
	"sync"

	"github.com/google/uuid"
)

type MemoryRepoProvider struct {
	lockStoredData sync.RWMutex
	storedData     map[uuid.UUID][]byte
}

func NewMemoryRepoProvider() *MemoryRepoProvider {

	return &MemoryRepoProvider{
		storedData: map[uuid.UUID][]byte{},
	}
}

func (repo *MemoryRepoProvider) Write(ID uuid.UUID, data []byte) error {

	repo.lockStoredData.Lock()
	defer repo.lockStoredData.Unlock()

	if _, ok := repo.storedData[ID]; ok {
		// TODO yep
		panic("writing twice the same flowfile ")
	}

	repo.storedData[ID] = data

	return nil
}

func (repo *MemoryRepoProvider) Remove(ID uuid.UUID) error {

	repo.lockStoredData.Lock()
	defer repo.lockStoredData.Unlock()

	if _, ok := repo.storedData[ID]; !ok {
		// TODO yep
		panic("Removing twice the same flowfile ")
	}

	delete(repo.storedData, ID)

	return nil
}
