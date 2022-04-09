package messaging

import (
	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/storage"
)

type Task struct {
	Recipient   types.Recipient `json:"recipient"`
	Destination types.Platform  `json:"destination"`
}

type Alerts struct {
	tasks []types.AlertTask
}

func NewAlerts() *Alerts {
	return &Alerts{tasks: make([]types.AlertTask, 0)}
}

func (a *Alerts) LoadAlerts(url string) []Task {
	var tasks []Task
	storage.Get(url, &tasks)
	return tasks
}

func (a *Alerts) AddAlert(url string, task Task) {
	tasks := a.LoadAlerts(url)
	//check if task already exists
	for _, t := range tasks {
		if t.Recipient == task.Recipient && t.Destination == task.Destination {
			return
		}
	}
	tasks = append(tasks, task)
	a.SaveAlerts(url, tasks)
}

func (a *Alerts) SaveAlerts(url string, tasks []Task) {
	storage.Set(url, tasks, 0)
}

func (a *Alerts) DeleteTask(urls []string, recipient types.Recipient, platform types.Platform) int {
	numberOfDeletedTasks := 0
	for _, url := range urls {
		tasks := a.LoadAlerts(url)
		for i, task := range tasks {
			if task.Recipient == recipient && task.Destination == platform {
				tasks = append(tasks[:i], tasks[i+1:]...)
				numberOfDeletedTasks += 1
			}
		}
		a.SaveAlerts(url, tasks)
	}
	return numberOfDeletedTasks
}
