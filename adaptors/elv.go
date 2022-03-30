package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type ELV struct {
	c *colly.Collector
}

func NewELV(c *colly.Collector) *ELV {
	copy := c.Clone()
	return &ELV{copy}
}

func (b *ELV) Run(list utils.Websites) {
	for _, url := range list.GetUrls("elv") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://de.elv.com"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting ELV ", r.URL.String())
	// })

	b.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	b.c.OnHTML("h1.product--title", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.Ctx.Get("url"))
		item.Name = e.Text
		item.InStock = e.ChildText("#buy-button") == "In den Warenkorb"
		item.PriceString = e.ChildText(".price--content")
		item.Time = time.Now().Format("15:04:05")
		item.UnixTime = time.Now().Unix()
		list.UpdateItemInList(item)
	})

	b.c.OnHTML(".product--buybox.block", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.Ctx.Get("url"))
		item.PriceString = e.ChildText(".price--content.content--default")
		item.InStock = !(strings.Contains(e.ChildText(".delivery--text"), "nicht lieferbar"))
		list.UpdateItemInList(item)
	})

	b.c.OnHTML(".product--buybox.block", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.Ctx.Get("url"))
		item.InStock = e.ChildText(".delivery--text.delivery--text-available") != ""
		list.UpdateItemInList(item)
	})

}

func (b *ELV) Wait() {
	b.c.Wait()
}
