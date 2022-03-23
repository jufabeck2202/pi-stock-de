package utils

import (
	"log"
	"os"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"

	"github.com/jufabeck2202/piScraper/storage"
)

type Website struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Shop          string    `json:"shop"`
	URL           string    `json:"url"`
	Type          string    `json:"type"`
	Ram           int       `json:"ram"`
	InStock       bool      `json:"in_stock"`
	PriceString   string    `json:"price_string"`
	StockNumer    int       `json:"stock_number"`
	Time          string    `json:"time"`
	UpdateCounter int       `json:"update_counter"`
}

type Scrape struct {
	Websites []struct {
		URL  string `yaml:"url"`
		Type string `yaml:"type"`
		Ram  int    `yaml:"ram"`
		Shop string `yaml:"shop"`
	} `yaml:"websites"`
}
type Websites struct {
	list []Website
}

func (p Websites) GetList() []Website {
	return p.list
}
func (p Websites) Init() {
	p.list = make([]Website, 0)
	p.Load()
	data, err := os.ReadFile("./website.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	websites := Scrape{}
	err = yaml.Unmarshal([]byte(data), &websites)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, v := range websites.Websites {
		if !Contains(p.list, v.URL) {
			p.list = append(p.list, Website{URL: v.URL, Shop: v.Shop, Type: v.Type, Ram: v.Ram})
		}
	}
	p.Save()

}

func (p *Websites) Load() {
	err := storage.Get("websites", &p.list)
	if err != nil {
		log.Println("error loading: ", err)
	}
}
func (p *Websites) Save() {
	err := storage.Set("websites", p.list)
	if err != nil {
		log.Println("error saving: ", err)
	}
}

func (p Websites) UpdateItemInList(item Website) {
	for i, v := range p.list {
		if v.URL == item.URL {
			item.UpdateCounter = v.UpdateCounter + 1
			p.list[i] = item
		}
	}
}

//get PiItem from PiList by id
func (p *Websites) GetItemById(url string) Website {
	for _, v := range p.list {
		if v.URL == url {
			return v
		}
	}
	return Website{URL: url, Id: uuid.New()}
}

// get list of urls for shop
func (p *Websites) GetUrls(shop string) []string {
	var urls []string
	for _, v := range p.list {
		if v.Shop == shop {
			urls = append(urls, v.URL)
		}
	}
	return urls
}

func (p *Websites) CheckForChanges() []Website {
	updatedValues := make([]Website, 0)
	oldPiList := &Websites{}
	oldPiList.Load()
	for _, v := range p.list {
		if Contains(oldPiList.list, v.URL) {
			item := GetByUrl(oldPiList.list, v.URL)
			if item.InStock != v.InStock {
				if v.InStock {
					updatedValues = append(updatedValues, v)
				}
				// updatedValues = append(updatedValues, v)
			}
		}
	}

	return updatedValues
}
