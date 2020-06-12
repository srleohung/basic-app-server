package syncer

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Downloader interface {
	Download(resourceRelativePath string, cachePath string) (string, error)
	DownloadWithQueries(resourceRelativePath string, cachePath string, queries map[string]string) (string, error)
}

type HTTPDownloader struct {
	httpClient           http.Client
	scheme               string
	host                 string
	username             string
	password             string
	resourcesUrlTemplate url.URL
}

func NewHTTPDownloader(scheme, host, username, password string) *HTTPDownloader {
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
	return &HTTPDownloader{
		httpClient:           httpClient,
		scheme:               scheme,
		host:                 host,
		username:             username,
		password:             password,
		resourcesUrlTemplate: resourcesUrlTemplate,
	}
}

func (downloader *HTTPDownloader) Download(resourceRelativePath string, cachePath string) (string, error) {
	resourcesUrl := downloader.resourcesUrlTemplate
	resourcesUrl.Path = resourceRelativePath
	return downloader.parallelDownload(resourceRelativePath, cachePath, resourcesUrl)
}

func (downloader *HTTPDownloader) DownloadWithQueries(resourceRelativePath string, cachePath string, queries map[string]string) (string, error) {
	resourcesUrl := downloader.resourcesUrlTemplate
	query := url.Values{}
	for k, v := range queries {
		query.Set(k, v)
	}
	resourcesUrl.RawQuery = query.Encode()
	resourcesUrl.Path = resourceRelativePath
	return downloader.parallelDownload(resourceRelativePath, cachePath, resourcesUrl)
}

func (downloader *HTTPDownloader) parallelDownload(resourceRelativePath string, cachePath string, resourcesUrl url.URL) (string, error) {
	downloadCacheFilePath := path.Join(cachePath, resourceRelativePath)
	var err error
	if _, err = os.Stat(path.Dir(downloadCacheFilePath)); os.IsNotExist(err) {
		if err = os.MkdirAll(path.Dir(downloadCacheFilePath), 0775); err != nil {
			return "", err
		}
	}
	var wg sync.WaitGroup
	type partialFile struct {
		path   string
		min    int
		max    int
		size   int
		finish bool
	}
	res, err := http.Head(resourcesUrl.String())
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New(res.Status)
	}
	filename := path.Base(res.Request.URL.Path)
	maps := res.Header
	length, _ := strconv.Atoi(maps["Content-Length"][0]) // Get the content length from the header request]
	limit := 20                                          // 10 Go-routines for the process so each downloads 18.7MB
	lenSub := length / limit                             // Bytes for each Go-routine
	diff := length % limit                               // Get the remaining for the last request
	body := make([]string, limit)                        // Make up a temporary array to hold the data to be written to the file
	partial := make([]partialFile, limit)
	downloaded := true
	for i := range partial {
		min := lenSub * i       // Min range
		max := lenSub * (i + 1) // Max range
		if i == limit-1 {
			max += diff // Add the remaining bytes in the last request
			partial[i].size = lenSub + diff
		} else {
			partial[i].size = lenSub
		}
		partial[i].max = max
		partial[i].min = min
	}
	for index := range partial {
		wg.Add(1)
		go func(file *partialFile, downloaded bool, i int) {
			tmpFile := path.Join(cachePath, filename+"."+strconv.Itoa(i))
			file.path = tmpFile
			if info, err := os.Stat(tmpFile); os.IsNotExist(err) || info.Size() != int64(file.size) {
				client := &http.Client{}
				req, _ := http.NewRequest(http.MethodGet, resourcesUrl.String(), nil)
				req.SetBasicAuth(downloader.username, downloader.password)
				rangeHeader := "bytes=" + strconv.Itoa(file.min) + "-" + strconv.Itoa(file.max-1) // Add the data for the Range header of the form "bytes=0-100"
				req.Header.Add("Range", rangeHeader)
				resp, err := client.Do(req)
				if err != nil {
					return
				}
				defer resp.Body.Close()
				reader, _ := ioutil.ReadAll(resp.Body)
				body[i] = string(reader)
				failToWrite := ioutil.WriteFile(tmpFile, []byte(string(body[i])), 0777) // Write to the file i as a byte array
				if failToWrite == nil {
					retrieved, _ := ioutil.ReadFile(tmpFile)
					if len(retrieved) != file.size {
						downloaded = false
					}
				}
			}
			wg.Done()
		}(&partial[index], downloaded, index)
	}
	wg.Wait()
	if downloaded {
		var whole string
		for _, element := range partial {
			reader, _ := ioutil.ReadFile(element.path)
			whole += string(reader)
			os.Remove(element.path)
		}
		if err = ioutil.WriteFile(path.Join(cachePath, resourceRelativePath), []byte(whole), 0777); err != nil {
			return "", err
		}
		if _, err = os.Stat(path.Join(cachePath, resourceRelativePath)); err != nil {
			return "", err
		} else {
			return path.Join(cachePath, resourceRelativePath), nil
		}
	} else {
		return "", err
	}
}
