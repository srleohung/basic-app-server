package main

import (
	"basic-app-server/config"
	"basic-app-server/core"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_appServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appServer()
		})
	}
}

func Test_watchdog(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			watchdog()
		})
	}
}

func Test_loadConfig(t *testing.T) {
	tests := []struct {
		name string
		want *config.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appServerSetup(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want *core.AppServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appServerSetup(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appServerSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}
