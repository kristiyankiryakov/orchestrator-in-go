package manager

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/kristiyankiryakov/orchestrator-in-go/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}
