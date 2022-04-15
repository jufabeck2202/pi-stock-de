package alertsrv

import (
	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
	"github.com/jufabeck2202/piScraper/messaging/types"
)

type service struct {
	tasks           []types.AlertTask
	alertRepository ports.RedisRepository
}

func New(alertRepository ports.RedisRepository) *service {
	return &service{
		alertRepository: alertRepository,
		tasks:           make([]types.AlertTask, 0),
	}
}

func (srv *service) LoadAlerts(url string) []domain.Alert {
	var tasks []domain.Alert
	//TODO add error handling
	srv.alertRepository.Get(url, &tasks)
	return tasks
}

func (srv *service) AddAlert(url string, task domain.Alert) {
	//TODO add error handling
	tasks := srv.LoadAlerts(url)
	//check if task already exists
	for _, t := range tasks {
		if t.Recipient == task.Recipient && t.Destination == task.Destination {
			return
		}
	}
	tasks = append(tasks, task)
	srv.SaveAlerts(url, tasks)
}

func (srv *service) SaveAlerts(url string, tasks []domain.Alert) {
	srv.alertRepository.Set(url, tasks, 0)
}

func (srv *service) DeleteTask(urls []string, recipient types.Recipient, platform types.Platform) int {
	numberOfDeletedTasks := 0
	for _, url := range urls {
		tasks := srv.LoadAlerts(url)
		for i, task := range tasks {
			if task.Recipient == recipient && task.Destination == platform {
				tasks = append(tasks[:i], tasks[i+1:]...)
				numberOfDeletedTasks += 1
			}
		}
		srv.SaveAlerts(url, tasks)
	}
	return numberOfDeletedTasks
}
