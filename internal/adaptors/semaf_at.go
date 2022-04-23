package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"
	"github.com/google/uuid"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type Semaf struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewSemaf(c *colly.Collector, websiteService ports.WebsiteService) *Semaf {
	copy := c.Clone()
	return &Semaf{copy, websiteService}
}

func (s *Semaf) Run() {
	for _, url := range s.websiteService.GetUrls("semaf") {
		s.c.Visit(url)
	}
	s.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://electronics.semaf.at/"),
	}
	// s.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Semaf: ", r.URL.String())
	// })
	s.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	s.c.OnHTML(".product-offer", func(e *colly.HTMLElement) {
		item := s.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Id = uuid.New()
		item.InStock = e.ChildText(".signal_image") != "Nicht auf Lager"
		item.PriceString = e.ChildText(".price")
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		s.websiteService.UpdateItemInList(item)
	})

	s.c.OnHTML(".product-headline", func(e *colly.HTMLElement) {
		item := s.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText(".product-title")
		s.websiteService.UpdateItemInList(item)
	})
}

func (s *Semaf) Wait() {
	s.c.Wait()
}
