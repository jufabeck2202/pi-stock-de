package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type Okdo struct {
	c              *colly.Collector
	websiteService ports.WebsiteService
}

func NewOkdo(c *colly.Collector, websiteService ports.WebsiteService) *Okdo {
	copy := c.Clone()
	return &Okdo{copy, websiteService}
}

func (b *Okdo) Run() {
	for _, url := range b.websiteService.GetUrls("okdo") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.okdo.com/de/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Okdo ", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML(".single-product-summary", func(e *colly.HTMLElement) {
		item := b.websiteService.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText(".c-product__title")
		item.InStock = !(e.ChildText(".c-stock-level.c-stock-level--low") == "Ausverkauft")
		item.PriceString = strings.Split(e.ChildText(".woocommerce-Price-amount.amount"), "€")[1] + " €"
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		b.websiteService.UpdateItemInList(item)
	})

}

func (b *Okdo) Wait() {
	b.c.Wait()
}
