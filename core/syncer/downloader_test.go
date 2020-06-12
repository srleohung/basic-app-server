package syncer

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewHTTPDownloader(t *testing.T) {
	type args struct {
		scheme   string
		host     string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *HTTPDownloader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPDownloader(tt.args.scheme, tt.args.host, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPDownloader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPDownloader_Download(t *testing.T) {
	type args struct {
		resourceRelativePath string
		cachePath            string
	}
	tests := []struct {
		name       string
		downloader *HTTPDownloader
		args       args
		want       string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.downloader.Download(tt.args.resourceRelativePath, tt.args.cachePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPDownloader.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPDownloader.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPDownloader_DownloadWithQueries(t *testing.T) {
	type args struct {
		resourceRelativePath string
		cachePath            string
		queries              map[string]string
	}
	tests := []struct {
		name       string
		downloader *HTTPDownloader
		args       args
		want       string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.downloader.DownloadWithQueries(tt.args.resourceRelativePath, tt.args.cachePath, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPDownloader.DownloadWithQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPDownloader.DownloadWithQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPDownloader_parallelDownload(t *testing.T) {
	type args struct {
		resourceRelativePath string
		cachePath            string
		resourcesUrl         url.URL
	}
	tests := []struct {
		name       string
		downloader *HTTPDownloader
		args       args
		want       string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.downloader.parallelDownload(tt.args.resourceRelativePath, tt.args.cachePath, tt.args.resourcesUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPDownloader.parallelDownload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPDownloader.parallelDownload() = %v, want %v", got, tt.want)
			}
		})
	}
}
