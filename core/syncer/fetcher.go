package syncer

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Fetcher interface {
	Fetch(resourceRelativePath string) ([]byte, error)
	FetchWithQueries(resourceRelativePath string, queries map[string]string) ([]byte, error)
}

type HTTPFetcher struct {
	httpClient           http.Client
	scheme               string
	host                 string
	username             string
	password             string
	resourcesUrlTemplate url.URL
}

func NewHTTPFetcher(scheme, host, username, password string) *HTTPFetcher {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 30 * time.Second,
	}
	httpClient := http.Client{Transport: transport}
	resourcesUrlTemplate := url.URL{Scheme: scheme, Host: host}
	return &HTTPFetcher{
		httpClient:           httpClient,
		scheme:               scheme,
		host:                 host,
		username:             username,
		password:             password,
		resourcesUrlTemplate: resourcesUrlTemplate,
	}
}

func (fetcher *HTTPFetcher) Fetch(resourceRelativePath string) ([]byte, error) {
	resourcesUrl := fetcher.resourcesUrlTemplate
	resourcesUrl.Path = resourceRelativePath
	return fetcher.fetchingUrl(resourcesUrl)
}

func (fetcher *HTTPFetcher) FetchWithQueries(resourceRelativePath string, queries map[string]string) ([]byte, error) {
	resourcesUrl := fetcher.resourcesUrlTemplate
	query := url.Values{}
	for k, v := range queries {
		query.Set(k, v)
	}
	resourcesUrl.RawQuery = query.Encode()
	resourcesUrl.Path = resourceRelativePath
	return fetcher.fetchingUrl(resourcesUrl)
}

func (fetcher *HTTPFetcher) fetchingUrl(resourcesUrl url.URL) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, resourcesUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(fetcher.username, fetcher.password)
	res, err := fetcher.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK:
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	case http.StatusNotModified:
		return nil, errors.New(res.Status)
	case http.StatusTeapot:
		return nil, errors.New(res.Status)
	default:
		return nil, errors.New(res.Status)
	}
}
