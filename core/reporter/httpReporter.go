package reporter

import (
	"bytes"
	"compress/gzip"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type Reporter interface {
	Report(reportPath string, buffer *bytes.Buffer) error
	ReportWithQueries(reportPath string, buffer *bytes.Buffer, queries map[string]string) (bool, error)
}

type HTTPReporter struct {
	httpClient           http.Client
	scheme               string
	host                 string
	username             string
	password             string
	reportingUrlTemplate url.URL
}

func NewHTTPReporter(scheme, host, username, password string) *HTTPReporter {
	httpClient := http.Client{Timeout: 2 * time.Minute}

	reportingUrlTemplate := url.URL{Scheme: scheme, Host: host}

	return &HTTPReporter{
		httpClient:           httpClient,
		scheme:               scheme,
		host:                 host,
		username:             username,
		password:             password,
		reportingUrlTemplate: reportingUrlTemplate,
	}
}

func (reporter *HTTPReporter) Report(reportPath string, buffer *bytes.Buffer) error {
	reportingUrl := reporter.reportingUrlTemplate
	reportingUrl.Path = reportPath
	if req, err := reporter.formReportingUrl(reportingUrl, buffer); err != nil {
		return err
	} else {
		res, err := reporter.httpClient.Do(req)
		if err != nil {
			return err
		} else {
			if res.StatusCode != http.StatusOK {
				return errors.New(res.Status)
			}
		}
		return nil
	}
}

func (reporter *HTTPReporter) ReportWithQueries(reportPath string, buffer *bytes.Buffer, queries map[string]string) (bool, error) {
	reportingUrl := reporter.reportingUrlTemplate
	query := url.Values{}
	for k, v := range queries {
		query.Set(k, v)
	}
	reportingUrl.RawQuery = query.Encode()
	reportingUrl.Path = reportPath
	if req, err := reporter.formReportingUrl(reportingUrl, buffer); err != nil {
		return false, err
	} else {
		res, err := reporter.httpClient.Do(req)
		if err != nil {
			return false, err
		} else {
			if res.StatusCode != http.StatusOK {
				if res.StatusCode == 307 {
					return true, errors.New(res.Status)
				}
				return false, errors.New(res.Status)
			}
		}
		defer res.Body.Close()
		return false, nil
	}
}

func (reporter *HTTPReporter) formReportingUrl(reportingUrl url.URL, buffer *bytes.Buffer) (*http.Request, error) {
	// try gzip compression when ever possible
	var gziperr error = nil
	gzipBuffer := new(bytes.Buffer)
	if zw, err := gzip.NewWriterLevel(gzipBuffer, gzip.BestCompression); err != nil {
		gziperr = err
	} else if _, err := zw.Write(buffer.Bytes()); err != nil {
		gziperr = err
	} else if err := zw.Flush(); err != nil {
		gziperr = err
	} else if err := zw.Close(); err != nil {
		gziperr = err
	}

	if gziperr != nil {
		if req, err := http.NewRequest(http.MethodPost, reportingUrl.String(), buffer); err != nil {
			return nil, err
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(reporter.username, reporter.password)
			return req, nil
		}
	} else {
		if req, err := http.NewRequest(http.MethodPost, reportingUrl.String(), gzipBuffer); err != nil {
			return nil, err
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Content-Encoding", "gzip")
			req.SetBasicAuth(reporter.username, reporter.password)
			return req, nil
		}
	}
}
