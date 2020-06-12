package manager

import (
	"basic-app-server/datastore"
	"basic-app-server/types"
	"reflect"
	"testing"
)

func TestNewExampleManager(t *testing.T) {
	type args struct {
		datastore datastore.Datastore
	}
	tests := []struct {
		name string
		args args
		want *ExampleManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExampleManager(tt.args.datastore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExampleManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleManager_Select(t *testing.T) {
	tests := []struct {
		name    string
		manager *ExampleManager
		want    []types.ExampleData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.manager.Select(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExampleManager.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleManager_InsertOne(t *testing.T) {
	type args struct {
		exampleData types.ExampleData
	}
	tests := []struct {
		name    string
		manager *ExampleManager
		args    args
		want    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.manager.InsertOne(tt.args.exampleData); got != tt.want {
				t.Errorf("ExampleManager.InsertOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
