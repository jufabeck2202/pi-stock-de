package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type BuyZero struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewBuyZero(c *colly.Collector, websiteService ports.WebsiteService) *BuyZero {
	copy := c.Clone()
	return &BuyZero{copy, websiteService}
}

func (s *BuyZero) Run() {
	for _, url := range s.websiteService.GetUrls("buyzero") {
		s.c.Visit(url)
	}
	s.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://buyzero.de/"),
	}
	// s.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Buyzero: ", r.URL.String())
	// })

	s.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	s.c.OnHTML(".product-meta__title.heading.h1", func(e *colly.HTMLElement) {
		item := s.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.Text
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		s.websiteService.UpdateItemInList(item)
	})

	s.c.OnHTML(".price-list", func(e *colly.HTMLElement) {
		item := s.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.PriceString = strings.Replace(e.ChildText(".price"), "Angebotspreis", "", -1)
		s.websiteService.UpdateItemInList(item)
	})

	s.c.OnHTML(".product-form__inventory.inventory", func(e *colly.HTMLElement) {
		item := s.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.InStock = e.Text != "Ausverkauft"
		s.websiteService.UpdateItemInList(item)
	})
}

func (s *BuyZero) Wait() {
	s.c.Wait()
}
