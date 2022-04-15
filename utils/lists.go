package utils

import "github.com/jufabeck2202/piScraper/internal/core/domain"

func Contains(list []domain.Website, url string) bool {
	for _, site := range list {
		if site.URL == url {
			return true
		}
	}
	return false
}

func GetByUrl(list []domain.Website, url string) domain.Website {
	for _, site := range list {
		if site.URL == url {
			return site
		}
	}
	return domain.Website{}
}
