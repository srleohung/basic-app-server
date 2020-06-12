package tasks

import (
	"basic-app-server/config"
	"basic-app-server/core/manager"
	"basic-app-server/core/syncer"
	"reflect"
	"sync"
	"testing"
)

func TestNewExampleFetchingTask(t *testing.T) {
	type args struct {
		manager   *manager.ExampleManager
		syncer    syncer.Syncer
		config    config.ExampleFetchingTaskConfig
		taskMutex *sync.Mutex
	}
	tests := []struct {
		name string
		args args
		want *ExampleFetchingTask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExampleFetchingTask(tt.args.manager, tt.args.syncer, tt.args.config, tt.args.taskMutex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExampleFetchingTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleFetchingTask_Start(t *testing.T) {
	tests := []struct {
		name    string
		task    *ExampleFetchingTask
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.task.Start(); (err != nil) != tt.wantErr {
				t.Errorf("ExampleFetchingTask.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExampleFetchingTask_fetch(t *testing.T) {
	tests := []struct {
		name string
		task *ExampleFetchingTask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.task.fetch()
		})
	}
}

func TestExampleFetchingTask_Stop(t *testing.T) {
	tests := []struct {
		name    string
		task    *ExampleFetchingTask
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.task.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("ExampleFetchingTask.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExampleFetchingTask_IsRunning(t *testing.T) {
	tests := []struct {
		name string
		task *ExampleFetchingTask
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.IsRunning(); got != tt.want {
				t.Errorf("ExampleFetchingTask.IsRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleFetchingTask_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		task *ExampleFetchingTask
		want TaskStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExampleFetchingTask.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
