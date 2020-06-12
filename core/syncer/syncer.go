package syncer

import (
	"basic-app-server/logger"
	"path"

	"github.com/sirupsen/logrus"
)

var syncerLogger *logrus.Entry = logger.GetLogger("syncer")

type Syncer interface {
	Fetch(resourceRelativePath string) ([]byte, error)
	FetchWithQueries(resourceRelativePath string, queries map[string]string) ([]byte, error)
	Download(resourceRelativePath string, cachePath string) (string, error)
	DownloadWithQueries(resourceRelativePath string, cachePath string, queries map[string]string) (string, error)
}

type HTTPSyncer struct {
	fetcher           Fetcher
	downloader        Downloader
	localResourcePath string
	cachePath         string
}

func NewHTTPSyncer(localResourcePath string, scheme, host, username, password string) *HTTPSyncer {
	return &HTTPSyncer{
		fetcher:           NewHTTPFetcher(scheme, host, username, password),
		downloader:        NewHTTPDownloader(scheme, host, username, password),
		localResourcePath: localResourcePath,
		cachePath:         path.Join(localResourcePath, ".app_cache"),
	}
}

func (syncer *HTTPSyncer) Fetch(resourceRelativePath string) ([]byte, error) {
	return syncer.fetcher.Fetch(resourceRelativePath)
}

func (syncer *HTTPSyncer) FetchWithQueries(resourceRelativePath string, queries map[string]string) ([]byte, error) {
	return syncer.fetcher.FetchWithQueries(resourceRelativePath, queries)
}

func (syncer *HTTPSyncer) Download(resourceRelativePath string, cachePath string) (string, error) {
	return syncer.downloader.Download(resourceRelativePath, cachePath)
}

func (syncer *HTTPSyncer) DownloadWithQueries(resourceRelativePath string, cachePath string, queries map[string]string) (string, error) {
	return syncer.downloader.DownloadWithQueries(resourceRelativePath, cachePath, queries)
}
