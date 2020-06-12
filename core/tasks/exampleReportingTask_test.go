package tasks

import (
	"basic-app-server/config"
	"basic-app-server/core/manager"
	"basic-app-server/core/reporter"
	"reflect"
	"testing"
)

func TestNewExampleReportingTask(t *testing.T) {
	type args struct {
		manager  *manager.ExampleManager
		reporter reporter.Reporter
		config   config.ExampleReportingTaskConfig
	}
	tests := []struct {
		name string
		args args
		want *ExampleReportingTask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExampleReportingTask(tt.args.manager, tt.args.reporter, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExampleReportingTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleReportingTask_Start(t *testing.T) {
	tests := []struct {
		name    string
		task    *ExampleReportingTask
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.task.Start(); (err != nil) != tt.wantErr {
				t.Errorf("ExampleReportingTask.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExampleReportingTask_report(t *testing.T) {
	tests := []struct {
		name string
		task *ExampleReportingTask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.task.report()
		})
	}
}

func TestExampleReportingTask_Stop(t *testing.T) {
	tests := []struct {
		name    string
		task    *ExampleReportingTask
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.task.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("ExampleReportingTask.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExampleReportingTask_IsRunning(t *testing.T) {
	tests := []struct {
		name string
		task *ExampleReportingTask
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.IsRunning(); got != tt.want {
				t.Errorf("ExampleReportingTask.IsRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleReportingTask_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		task *ExampleReportingTask
		want TaskStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExampleReportingTask.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
