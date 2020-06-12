package types

import (
	"reflect"
	"testing"
)

func TestExampleData_Format(t *testing.T) {
	tests := []struct {
		name        string
		exampleData *ExampleData
		want        ExampleData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exampleData.Format(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExampleData.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleData_Init(t *testing.T) {
	tests := []struct {
		name        string
		exampleData *ExampleData
		want        ExampleData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exampleData.Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExampleData.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
