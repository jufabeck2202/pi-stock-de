package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type Reichelt struct {
	c *colly.Collector
}

func NewReichelt(c *colly.Collector) *Reichelt {
	copy := c.Clone()
	return &Reichelt{copy}
}

func (b *Reichelt) Run(list utils.Websites) {
	for _, url := range list.GetUrls("reichelt") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.reichelt.de/"),
	}
	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML("#article", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText("#av_articleheader")
		item.InStock = !(e.ChildText(".availability") == "z.Zt. ausverkauft")
		item.PriceString = e.ChildText("#av_price") + " €"
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

}

func (b *Reichelt) Wait() {
	b.c.Wait()
}
