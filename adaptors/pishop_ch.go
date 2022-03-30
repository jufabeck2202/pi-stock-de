package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type Pishop struct {
	c *colly.Collector
}

func NewPishop(c *colly.Collector) *Pishop {
	copy := c.Clone()
	return &Pishop{copy}
}

func (b *Pishop) Run(list utils.Websites) {
	for _, url := range list.GetUrls("pishopch") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.pi-shop.ch/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Pishop ", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML(".product-primary-column.product-shop", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.ChildText(".product-name")
		item.InStock = !(e.ChildText(".availability.out-of-stock") == "Verfügbarkeit: Zur Zeit nicht an Lager")
		item.Time = time.Now().Format("15:04:05")
		item.PriceString = e.ChildText(".price")
		item.UnixTime = time.Now().Unix()
		list.UpdateItemInList(item)
	})

}

func (b *Pishop) Wait() {
	b.c.Wait()
}
