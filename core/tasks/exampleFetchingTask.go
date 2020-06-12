package tasks

import (
	"basic-app-server/config"
	"basic-app-server/core/manager"
	"basic-app-server/core/syncer"
	"basic-app-server/logger"
	"basic-app-server/types"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var exampleFetchingTaskLogger *logrus.Entry = logger.GetLogger("exampleFetchingTask")
var resourceRelativePath = "/mclient/plan"

type ExampleFetchingTask struct {
	running      bool
	manager      *manager.ExampleManager
	syncer       syncer.Syncer
	taskMutex    *sync.Mutex
	period       int
	stopSyncChan chan bool
}

func NewExampleFetchingTask(manager *manager.ExampleManager, syncer syncer.Syncer, config config.ExampleFetchingTaskConfig, taskMutex *sync.Mutex) *ExampleFetchingTask {
	return &ExampleFetchingTask{
		running:      false,
		manager:      manager,
		syncer:       syncer,
		taskMutex:    taskMutex,
		period:       config.Period,
		stopSyncChan: make(chan bool, 1),
	}
}

func (task *ExampleFetchingTask) Start() error {
	if task.running {
		return nil
	}
	go func() {
		task.fetch()
		for {
			select {
			case <-task.stopSyncChan:
				break
			case <-time.After(time.Duration(task.period) * time.Second):
				task.fetch()
			}
		}
		exampleFetchingTaskLogger.Warnln("Fetching Task Stopped Running")
		task.running = false
	}()
	return nil
}

func (task *ExampleFetchingTask) fetch() {
	task.taskMutex.Lock()
	defer task.taskMutex.Unlock()
	queries := make(map[string]string)
	queries["code"] = "SmartRetail"
	bytes, err := task.syncer.FetchWithQueries(resourceRelativePath, queries)
	if err != nil {
		exampleFetchingTaskLogger.Warnln(err)
	} else {
		var exampleData types.ExampleData
		if err := json.Unmarshal(bytes, &exampleData); err != nil {
			exampleFetchingTaskLogger.Error(err)
		}
		exampleFetchingTaskLogger.Debugf("%#v", exampleData)
		if success := task.manager.InsertOne(exampleData); success {
			exampleFetchingTaskLogger.Info("Data successfully inserted")
		} else {
			exampleFetchingTaskLogger.Error("Failed to insert data")
		}
	}
}

func (task *ExampleFetchingTask) Stop() error {
	if !task.running {
		return nil
	}
	task.stopSyncChan <- true
	sleepCounter := 0
	for task.running {
		time.Sleep(2 * time.Second)
		sleepCounter += 2
		if sleepCounter > task.period+6 {
			return errors.New("Unable to stop reporting task")
		}
	}
	return nil
}

func (task *ExampleFetchingTask) IsRunning() bool {
	return task.running
}

func (task *ExampleFetchingTask) GetStatus() TaskStatus {
	if task.running {
		return TASK_STATUS_ENABLED
	} else {
		return TASK_STATUS_DISABLED
	}
}
