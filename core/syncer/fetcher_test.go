package syncer

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewHTTPFetcher(t *testing.T) {
	type args struct {
		scheme   string
		host     string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *HTTPFetcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPFetcher(tt.args.scheme, tt.args.host, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPFetcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPFetcher_Fetch(t *testing.T) {
	type args struct {
		resourceRelativePath string
	}
	tests := []struct {
		name    string
		fetcher *HTTPFetcher
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fetcher.Fetch(tt.args.resourceRelativePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPFetcher.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPFetcher.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPFetcher_FetchWithQueries(t *testing.T) {
	type args struct {
		resourceRelativePath string
		queries              map[string]string
	}
	tests := []struct {
		name    string
		fetcher *HTTPFetcher
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fetcher.FetchWithQueries(tt.args.resourceRelativePath, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPFetcher.FetchWithQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPFetcher.FetchWithQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPFetcher_fetchingUrl(t *testing.T) {
	type args struct {
		resourcesUrl url.URL
	}
	tests := []struct {
		name    string
		fetcher *HTTPFetcher
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fetcher.fetchingUrl(tt.args.resourcesUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPFetcher.fetchingUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPFetcher.fetchingUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
