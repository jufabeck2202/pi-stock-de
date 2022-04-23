package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/feeds"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type RssHandler struct {
	websiteService ports.WebsiteService
}

func NewRssHandler(websiteService ports.WebsiteService) *RssHandler {
	return &RssHandler{
		websiteService: websiteService,
	}
}

func (hdl *RssHandler) Get(c *fiber.Ctx) error {
	list := hdl.websiteService.GetList()
	items := []*feeds.Item{}

	feed := &feeds.Feed{
		Title:       "German Raspberry Pi Stock",
		Link:        &feeds.Link{Href: "https://pi.juli.sh"},
		Description: "Check if Raspberry Pis are in stock in german shops",
		Author:      &feeds.Author{Name: "Julian Beck", Email: "mail@julianbeck.com"},
		Created:     time.Now(),
	}
	for _, v := range list {
		if v.InStock {

			item := feeds.Item{
				Title:       v.Name,
				Link:        &feeds.Link{Href: v.URL},
				Description: v.PriceString,
			}
			items = append(items, &item)
		}
	}
	feed.Items = items
	c.Type("rss")
	rss, err := feed.ToRss()
	if err != nil {
		log.Println(err)
	}
	return c.SendString(rss)

}
