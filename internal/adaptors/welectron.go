package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type Welectron struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewWelectron(c *colly.Collector, websiteService ports.WebsiteService) *Welectron {
	copy := c.Clone()
	return &Welectron{copy, websiteService}
}

func (b *Welectron) Run() {
	for _, url := range b.websiteService.GetUrls("welectron") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.welectron.com/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Welectron ", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML(".product-info.col-sm-7", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText(".fn.product-title")
		item.InStock = !(e.ChildText(".status-text") == "nicht lieferbar") && !(e.ChildText(".status.status-1") == "im Zulauf")
		item.PriceString = strings.Split(e.ChildText(".prodprice.inclvat.text-nowrap"), "€")[0] + " €"
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		b.websiteService.UpdateItemInList(item)
	})

}

func (b *Welectron) Wait() {
	b.c.Wait()
}
