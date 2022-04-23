package domain

func Contains(list []Website, url string) bool {
	for _, site := range list {
		if site.URL == url {
			return true
		}
	}
	return false
}

func GetByUrl(list []Website, url string) Website {
	for _, site := range list {
		if site.URL == url {
			return site
		}
	}
	return Website{}
}
