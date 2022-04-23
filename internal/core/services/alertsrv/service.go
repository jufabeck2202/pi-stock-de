package alertsrv

import (
	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type service struct {
	tasks           []domain.AlertTask
	alertRepository ports.RedisRepository
	websiteService  ports.WebsiteService
}

func New(alertRepository ports.RedisRepository, websiteService ports.WebsiteService) *service {
	return &service{
		alertRepository: alertRepository,
		websiteService:  websiteService,
		tasks:           make([]domain.AlertTask, 0),
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

func (srv *service) DeleteTask(urls []string, recipient domain.Recipient, platform domain.Platform) int {
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

func (srv *service) RemoveEmailAlert(email string) int {
	//deletes all tasks for the given email

	numberOfDeletedTasks := 0
	for _, url := range srv.websiteService.GetAllUrls() {
		tasks := srv.LoadAlerts(url)
		for i, task := range tasks {
			if task.Recipient.Email == email {
				tasks = append(tasks[:i], tasks[i+1:]...)
				numberOfDeletedTasks += 1
			}
		}
		srv.SaveAlerts(url, tasks)
	}
	return numberOfDeletedTasks
}
