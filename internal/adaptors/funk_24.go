package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type Funk24 struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewFunk24(c *colly.Collector, websiteService ports.WebsiteService) *Funk24 {
	copy := c.Clone()
	return &Funk24{copy, websiteService}
}

func (b *Funk24) Run() {
	for _, url := range b.websiteService.GetUrls("funk24") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://shop.funk24.net/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("funk24", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML(".content-main--inner", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText("h1.product--title")
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		b.websiteService.UpdateItemInList(item)
	})

	b.c.OnHTML(".product--buybox.block", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.InStock = !(e.ChildText(".alert--content") == "Dieser Artikel steht derzeit nicht zur Verfügung!") && !(e.ChildText(".alert--content") == "Dieser Artikel steht derzeit nicht zur Verfügung!  Benachrichtigen Sie mich, sobald der Artikel lieferbar ist.")
		item.PriceString = strings.Split(e.ChildText(".price--content.content--default"), "€")[0] + " €"
		b.websiteService.UpdateItemInList(item)
	})
}

func (b *Funk24) Wait() {
	b.c.Wait()
}
