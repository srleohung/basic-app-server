package core

import (
	"basic-app-server/config"
	. "basic-app-server/core/manager"
	. "basic-app-server/core/reporter"
	. "basic-app-server/core/syncer"
	"basic-app-server/core/tasks"
	. "basic-app-server/core/tasks"
	. "basic-app-server/datastore"
	"basic-app-server/logger"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var coreLogger *logrus.Entry = logger.GetLogger("core")

var DEBUG_MODE bool = false

func init() {
	appServerEnv := os.Getenv("APP_SERVER_ENV")
	if strings.Contains(strings.ToUpper(appServerEnv), "DEBUG") {
		DEBUG_MODE = true
	}
}

type AppServer struct {
	config config.Config
	Datastore
	Syncer
	Reporter
	Tasks []tasks.Task
}

var appServerInstance *AppServer

func GetAppServer(configuration *config.Config, versionString string) *AppServer {
	if appServerInstance != nil {
		return appServerInstance
	}

	// Setup Datastore
	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Start setting up datastore.")
	var datastore Datastore
	switch configuration.DatastoreSettings.Type {
	default:
		datastore = NewBuiltinDatastore()
	}

	// Setup Syncer
	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Start setting up syncer.")
	syncer := NewHTTPSyncer(configuration.LocalResourcePath, configuration.WebServerSettings.Scheme, configuration.WebServerSettings.Host, configuration.WebServerSettings.Username, configuration.WebServerSettings.Password)

	// Setup Manager
	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Start setting up manager.")
	exampleManager := NewExampleManager(datastore)

	// Setup Reporter
	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Start setting up reporter.")
	reporter := NewHTTPReporter(configuration.WebServerSettings.Scheme, configuration.WebServerSettings.Host, configuration.WebServerSettings.Username, configuration.WebServerSettings.Password)

	// Setup Task
	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Start setting up task list.")
	task_list := []tasks.Task{}
	taskMutex := &sync.Mutex{}
	// Setup Reporting Task
	if reporterTask := NewExampleReportingTask(exampleManager, reporter, *configuration.Tasks.ExampleReportingTask); reporterTask != nil {
		task_list = append(task_list, reporterTask)
	}
	// Setup Fetching Task
	if fetchingTask := NewExampleFetchingTask(exampleManager, syncer, *configuration.Tasks.ExampleFetchingTask, taskMutex); fetchingTask != nil {
		task_list = append(task_list, fetchingTask)
	}

	coreLogger.Infoln("--------------------------------------------------")
	coreLogger.Infoln("Finished application server configuration and setup.")
	appServerInstance = &AppServer{
		config:    *configuration,
		Datastore: datastore,
		Syncer:    syncer,
		Reporter:  reporter,
		Tasks:     task_list,
	}
	return appServerInstance
}

func (appServer *AppServer) GetConfig() config.Config {
	return appServer.config
}

func (appServer *AppServer) Start() {
	for _, task := range appServer.Tasks {
		if err := task.Start(); err != nil {
			coreLogger.Error(err)
		}
		time.Sleep(1 * time.Second)
	}
}
