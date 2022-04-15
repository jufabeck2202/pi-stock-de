package websitesrv

import (
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
	"github.com/jufabeck2202/piScraper/utils"
)

type service struct {
	redisRepository ports.RedisRepository
	sites           domain.Websites
}

func New(websiteRepository ports.RedisRepository) *service {

	return &service{
		redisRepository: websiteRepository,
	}
}

func (srv *service) Init() {
	srv.sites = make(domain.Websites, 0)
	//load old data
	srv.Load()
	data, err := os.ReadFile("./website.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	websites := domain.WebsitesYaml{}
	err = yaml.Unmarshal([]byte(data), &websites)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, v := range websites.Websites {
		if !utils.Contains(srv.sites, v.URL) {
			srv.sites = append(srv.sites, domain.Website{URL: v.URL, Shop: v.Shop, Type: v.Type, Ram: v.Ram})
		}
	}
	srv.Save()

}

func (srv *service) Save() {
	err := srv.redisRepository.Set("websites", srv.sites, 0)
	if err != nil {
		log.Println("error saving: ", err)
	}
}

func (srv *service) Load() {
	err := srv.redisRepository.Get("websites", &srv.sites)
	if err != nil {
		log.Println("error loading: ", err)
	}
}

func (srv *service) GetList() domain.Websites {
	return srv.sites
}
func (srv *service) GetAllUrls() []string {
	var urls []string
	for _, v := range srv.sites {
		urls = append(urls, v.URL)
	}
	return urls
}

var mutex = &sync.Mutex{}

func (srv *service) UpdateItemInList(item domain.Website) {
	for i, v := range srv.sites {
		mutex.Lock()
		if v.URL == item.URL {
			item.UpdateCounter = v.UpdateCounter + 1
			srv.sites[i] = item
		}
		mutex.Unlock()
	}
}

//get PiItem from PiList by id
func (srv *service) GetItemById(url string) domain.Website {
	for _, v := range srv.sites {
		if v.URL == url {
			return v
		}
	}
	return domain.Website{URL: url, Id: uuid.New()}
}

// get list of urls for shop
func (srv *service) GetUrls(shop string) []string {
	var urls []string
	for _, v := range srv.sites {
		if v.Shop == shop {
			urls = append(urls, v.URL)
		}
	}
	return urls
}

func (srv *service) CheckForChanges() domain.Websites {
	updatedValues := make([]domain.Website, 0)
	oldPiList := New(srv.redisRepository)
	oldPiList.Load()
	for _, v := range srv.sites {
		if utils.Contains(oldPiList.sites, v.URL) {
			item := utils.GetByUrl(oldPiList.sites, v.URL)
			if item.InStock != v.InStock {
				if v.InStock {
					updatedValues = append(updatedValues, v)
				}
			}
		}
	}

	return updatedValues
}
