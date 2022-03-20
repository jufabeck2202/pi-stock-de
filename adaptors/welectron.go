package adaptors

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type Welectron struct {
	c *colly.Collector
}

func NewWelectron(c *colly.Collector) *Welectron {
	copy := c.Clone()
	return &Welectron{copy}
}

func (b *Welectron) Run(list utils.Websites) {
	for _, url := range list.GetUrls("welectron") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.welectron.com/"),
	}
	b.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Welectron ", r.URL.String())
	})

	b.c.OnHTML(".product-info.col-sm-7", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.Name = e.ChildText(".fn.product-title")
		item.InStock = !(e.ChildText(".status-text") == "nicht lieferbar") && !(e.ChildText(".status.status-1") == "im Zulauf")
		item.PriceString = strings.Split(e.ChildText(".prodprice.inclvat.text-nowrap"), "€")[0] + " €"
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

}

func (b *Welectron) Wait() {
	b.c.Wait()
}
