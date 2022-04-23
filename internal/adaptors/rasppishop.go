package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type Rasppishop struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewRappishop(c *colly.Collector, websiteService ports.WebsiteService) *Rasppishop {
	copy := c.Clone()
	return &Rasppishop{copy, websiteService}
}

func (b *Rasppishop) Run() {
	for _, url := range b.websiteService.GetUrls("rappishop") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.rasppishop.de/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Rasppishop ", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML("#result-wrapper", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText(".col-sm-10.col-md-6.col-lg-8")
		item.InStock = !(e.ChildText(".status.status-0") == "Produkt vergriffen")
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		b.websiteService.UpdateItemInList(item)
	})
	b.c.OnHTML(".product-info-box", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.PriceString = e.ChildText(".price_2")
		b.websiteService.UpdateItemInList(item)
	})
}

func (b *Rasppishop) Wait() {
	b.c.Wait()
}
