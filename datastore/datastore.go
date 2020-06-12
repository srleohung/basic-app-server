package datastore

import "basic-app-server/types"

type Datastore interface {
	Select() []types.ExampleData
	InsertOne(exampleData types.ExampleData) bool
}
