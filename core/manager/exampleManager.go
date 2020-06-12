package manager

import (
	"basic-app-server/datastore"
	"basic-app-server/types"
)

type ExampleManager struct {
	datastore datastore.Datastore
}

func NewExampleManager(datastore datastore.Datastore) *ExampleManager {
	return &ExampleManager{datastore: datastore}
}

func (manager *ExampleManager) Select() []types.ExampleData {
	return manager.datastore.Select()
}

func (manager *ExampleManager) InsertOne(exampleData types.ExampleData) bool {
	return manager.datastore.InsertOne(exampleData)
}
