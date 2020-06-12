package reporter

import (
	"bytes"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewHTTPReporter(t *testing.T) {
	type args struct {
		scheme   string
		host     string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *HTTPReporter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPReporter(tt.args.scheme, tt.args.host, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPReporter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPReporter_Report(t *testing.T) {
	type args struct {
		reportPath string
		buffer     *bytes.Buffer
	}
	tests := []struct {
		name     string
		reporter *HTTPReporter
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.reporter.Report(tt.args.reportPath, tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("HTTPReporter.Report() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPReporter_ReportWithQueries(t *testing.T) {
	type args struct {
		reportPath string
		buffer     *bytes.Buffer
		queries    map[string]string
	}
	tests := []struct {
		name     string
		reporter *HTTPReporter
		args     args
		want     bool
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.reporter.ReportWithQueries(tt.args.reportPath, tt.args.buffer, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPReporter.ReportWithQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPReporter.ReportWithQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPReporter_formReportingUrl(t *testing.T) {
	type args struct {
		reportingUrl url.URL
		buffer       *bytes.Buffer
	}
	tests := []struct {
		name     string
		reporter *HTTPReporter
		args     args
		want     *http.Request
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.reporter.formReportingUrl(tt.args.reportingUrl, tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPReporter.formReportingUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPReporter.formReportingUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
