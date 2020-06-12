package datastore

import (
	"basic-app-server/types"
	"sync"
)

type BuiltinDatastore struct {
	mutex       *sync.RWMutex
	exampleData []types.ExampleData
}

func NewBuiltinDatastore() *BuiltinDatastore {
	store := BuiltinDatastore{
		mutex:       &sync.RWMutex{},
		exampleData: []types.ExampleData{},
	}
	return &store
}

func (datastore *BuiltinDatastore) Select() []types.ExampleData {
	return datastore.exampleData
}

func (datastore *BuiltinDatastore) InsertOne(exampleData types.ExampleData) bool {
	datastore.exampleData = append(datastore.exampleData, exampleData)
	return true
}
