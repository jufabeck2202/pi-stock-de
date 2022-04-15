package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type Bechtle struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewBechtle(c *colly.Collector, websiteService ports.WebsiteService) *Bechtle {
	copy := c.Clone()
	return &Bechtle{copy, websiteService}
}

func (b *Bechtle) Run() {
	for _, url := range b.websiteService.GetUrls("bechtle") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.bechtle.com/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Bechtle ", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML(".organism.conversion-box.js-conversion-box.js-pds-conversion-box", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText("h1.h-h1.big-characters")
		item.InStock = !(e.ChildText(".delivery-info") == "Bestellen Sie jetzt und Sie erhalten die Ware sobald diese verfügbar ist.")
		item.PriceString = e.ChildText(".bechtle-price.js-price") + " €"
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		b.websiteService.UpdateItemInList(item)
	})

}

func (b *Bechtle) Wait() {
	b.c.Wait()
}
