package logger

import (
	"reflect"
	"testing"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Test_setLogDirectoryPath(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setLogDirectoryPath()
		})
	}
}

func Test_setPrimaryOutStream(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setPrimaryOutStream()
		})
	}
}

func Test_getPathMap(t *testing.T) {
	type args struct {
		logPath string
	}
	tests := []struct {
		name string
		args args
		want lfshook.PathMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPathMap(tt.args.logPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPathMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLogger(t *testing.T) {
	type args struct {
		module string
	}
	tests := []struct {
		name string
		args args
		want *logrus.Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogger(tt.args.module); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
