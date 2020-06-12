package datastore

import (
	"basic-app-server/types"
	"reflect"
	"testing"
)

func TestNewBuiltinDatastore(t *testing.T) {
	tests := []struct {
		name string
		want *BuiltinDatastore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuiltinDatastore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuiltinDatastore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuiltinDatastore_Select(t *testing.T) {
	tests := []struct {
		name      string
		datastore *BuiltinDatastore
		want      []types.ExampleData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.datastore.Select(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuiltinDatastore.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuiltinDatastore_InsertOne(t *testing.T) {
	type args struct {
		exampleData types.ExampleData
	}
	tests := []struct {
		name      string
		datastore *BuiltinDatastore
		args      args
		want      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.datastore.InsertOne(tt.args.exampleData); got != tt.want {
				t.Errorf("BuiltinDatastore.InsertOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
