package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type Bechtle struct {
	c *colly.Collector
}

func NewBechtle(c *colly.Collector) *Bechtle {
	copy := c.Clone()
	return &Bechtle{copy}
}

func (b *Bechtle) Run(list utils.Websites) {
	for _, url := range list.GetUrls("bechtle") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.bechtle.com/"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Bechtle ", r.URL.String())
	// })

	b.c.OnHTML(".organism.conversion-box.js-conversion-box.js-pds-conversion-box", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.Name = e.ChildText("h1.h-h1.big-characters")
		item.InStock = !(e.ChildText(".delivery-info") == "Bestellen Sie jetzt und Sie erhalten die Ware sobald diese verfügbar ist.")
		item.PriceString = e.ChildText(".bechtle-price.js-price") + " €"
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

}

func (b *Bechtle) Wait() {
	b.c.Wait()
}
