package syncer

import (
	"reflect"
	"testing"
)

func TestNewHTTPSyncer(t *testing.T) {
	type args struct {
		localResourcePath string
		scheme            string
		host              string
		username          string
		password          string
	}
	tests := []struct {
		name string
		args args
		want *HTTPSyncer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPSyncer(tt.args.localResourcePath, tt.args.scheme, tt.args.host, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPSyncer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPSyncer_Fetch(t *testing.T) {
	type args struct {
		resourceRelativePath string
	}
	tests := []struct {
		name    string
		syncer  *HTTPSyncer
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.syncer.Fetch(tt.args.resourceRelativePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSyncer.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPSyncer.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPSyncer_FetchWithQueries(t *testing.T) {
	type args struct {
		resourceRelativePath string
		queries              map[string]string
	}
	tests := []struct {
		name    string
		syncer  *HTTPSyncer
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.syncer.FetchWithQueries(tt.args.resourceRelativePath, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSyncer.FetchWithQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPSyncer.FetchWithQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPSyncer_Download(t *testing.T) {
	type args struct {
		resourceRelativePath string
		cachePath            string
	}
	tests := []struct {
		name    string
		syncer  *HTTPSyncer
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.syncer.Download(tt.args.resourceRelativePath, tt.args.cachePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSyncer.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPSyncer.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPSyncer_DownloadWithQueries(t *testing.T) {
	type args struct {
		resourceRelativePath string
		cachePath            string
		queries              map[string]string
	}
	tests := []struct {
		name    string
		syncer  *HTTPSyncer
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.syncer.DownloadWithQueries(tt.args.resourceRelativePath, tt.args.cachePath, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSyncer.DownloadWithQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPSyncer.DownloadWithQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
