package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type Okdo struct {
	c *colly.Collector
}

func NewOkdo(c *colly.Collector) *Okdo {
	copy := c.Clone()
	return &Okdo{copy}
}

func (b *Okdo) Run(list utils.Websites) {
	for _, url := range list.GetUrls("okdo") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.okdo.com/de/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Okdo ", r.URL.String())
	// })

	b.c.OnHTML(".single-product-summary", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.Name = e.ChildText(".c-product__title")
		item.InStock = !(e.ChildText(".c-stock-level.c-stock-level--low") == "Ausverkauft")
		item.PriceString = strings.Split(e.ChildText(".woocommerce-Price-amount.amount"), "€")[1] + " €"
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

}

func (b *Okdo) Wait() {
	b.c.Wait()
}
