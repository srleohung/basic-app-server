package tasks

import (
	"basic-app-server/config"
	"basic-app-server/core/manager"
	"basic-app-server/core/reporter"
	"basic-app-server/logger"
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

var exampleReportingTaskLogger *logrus.Entry = logger.GetLogger("exampleReportingTask")
var reportPath = "/mclient/plan"

type ExampleReportingTask struct {
	running      bool
	manager      *manager.ExampleManager
	reporter     reporter.Reporter
	period       int
	pageSize     int
	stopTaskChan chan bool
}

func NewExampleReportingTask(manager *manager.ExampleManager, reporter reporter.Reporter, config config.ExampleReportingTaskConfig) *ExampleReportingTask {
	return &ExampleReportingTask{
		manager:  manager,
		reporter: reporter,
		period:   config.Period,
		pageSize: config.PageSize,
		running:  false,
	}
}

func (task *ExampleReportingTask) Start() error {
	if task.running {
		return nil
	}
	go func() {
		for {
			select {
			case <-task.stopTaskChan:
				break
			case <-time.After(time.Duration(task.period) * time.Second):
				task.report()
			}
		}
		exampleReportingTaskLogger.Warnln("Report task has stopped")
		task.running = false
	}()
	return nil
}

func (task *ExampleReportingTask) report() {
	pageSize := task.pageSize
	data := task.manager.Select()
	exampleReportingTaskLogger.Infof("Report task started. Total %d.", len(data))
	for page, endIdx := 0, 0; endIdx != len(data); page++ {
		if (page+1)*pageSize > len(data) {
			endIdx = len(data)
		} else {
			endIdx = (page + 1) * pageSize
		}
		pending := data[page*pageSize : endIdx]
		buffer := new(bytes.Buffer)
		if err := json.NewEncoder(buffer).Encode(pending); err != nil {
			exampleReportingTaskLogger.Warnln(err)
			continue
		}
		if err := task.reporter.Report(reportPath, buffer); err != nil {
			exampleReportingTaskLogger.Warnln(err)
			continue
		}
	}
	exampleReportingTaskLogger.Infoln("Report task completion")
}

func (task *ExampleReportingTask) Stop() error {
	if !task.running {
		return nil
	}
	task.stopTaskChan <- true
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

func (task *ExampleReportingTask) IsRunning() bool {
	return task.running
}

func (task *ExampleReportingTask) GetStatus() TaskStatus {
	if task.running {
		return TASK_STATUS_ENABLED
	} else {
		return TASK_STATUS_DISABLED
	}
}
