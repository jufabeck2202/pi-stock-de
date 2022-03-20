package adaptors

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type Funk24 struct {
	c *colly.Collector
}

func NewFunk24(c *colly.Collector) *Funk24 {
	copy := c.Clone()
	return &Funk24{copy}
}

func (b *Funk24) Run(list utils.Websites) {
	for _, url := range list.GetUrls("funk24") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://shop.funk24.net/"),
	}
	b.c.OnRequest(func(r *colly.Request) {
		fmt.Println("funk24", r.URL.String())
	})

	b.c.OnHTML(".content-main--inner", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.Name = e.ChildText("h1.product--title")
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

	b.c.OnHTML(".product--buybox.block", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.InStock = !(e.ChildText(".alert--content") == "Dieser Artikel steht derzeit nicht zur Verfügung!") && !(e.ChildText(".alert--content") == "Dieser Artikel steht derzeit nicht zur Verfügung!  Benachrichtigen Sie mich, sobald der Artikel lieferbar ist.")
		item.PriceString = strings.Split(e.ChildText(".price--content.content--default"), "€")[0] + " €"
		list.UpdateItemInList(item)
	})
}

func (b *Funk24) Wait() {
	b.c.Wait()
}
