package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type BerryBase struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewBerryBase(c *colly.Collector, websiteService ports.WebsiteService) *BerryBase {
	copy := c.Clone()
	return &BerryBase{copy, websiteService}
}

func (b *BerryBase) Run() {
	for _, url := range b.websiteService.GetUrls("berrybase") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.berrybase.de"),
	}
	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML(".product--detail-upper", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText(".product--title")
		item.InStock = e.ChildText("#buy-button") == "In den Warenkorb"
		item.PriceString = e.ChildText(".price--content")
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		b.websiteService.UpdateItemInList(item)
	})

}

func (b *BerryBase) Wait() {
	b.c.Wait()
}
