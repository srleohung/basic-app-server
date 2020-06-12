package tasks

type TaskStatus string

const (
	TASK_STATUS_ENABLED  TaskStatus = "TASK_STATUS_ENABLED"
	TASK_STATUS_DISABLED TaskStatus = "TASK_STATUS_DISABLED"
)

type Task interface {
	Start() error
	Stop() error
	IsRunning() bool
	GetStatus() TaskStatus
}
