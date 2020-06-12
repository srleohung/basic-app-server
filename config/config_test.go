package config

import (
	"reflect"
	"testing"
)

func TestLoadConfigFromYaml(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfigFromYaml(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfigFromYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_SaveConfigToYamlFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		config  *Config
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.config.SaveConfigToYamlFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Config.SaveConfigToYamlFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
