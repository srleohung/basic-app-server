package core

import (
	"basic-app-server/config"
	"reflect"
	"testing"
)

func TestGetAppServer(t *testing.T) {
	type args struct {
		configuration *config.Config
		versionString string
	}
	tests := []struct {
		name string
		args args
		want *AppServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAppServer(tt.args.configuration, tt.args.versionString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAppServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppServer_GetConfig(t *testing.T) {
	tests := []struct {
		name      string
		appServer *AppServer
		want      config.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.appServer.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppServer.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppServer_Start(t *testing.T) {
	tests := []struct {
		name      string
		appServer *AppServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.appServer.Start()
		})
	}
}
