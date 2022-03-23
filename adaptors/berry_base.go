package adaptors

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"
	"github.com/google/uuid"

	"github.com/jufabeck2202/piScraper/utils"
)

type BerryBase struct {
	c *colly.Collector
}

func NewBerryBase(c *colly.Collector) *BerryBase {
	copy := c.Clone()
	return &BerryBase{copy}
}

func (b *BerryBase) Run(list utils.Websites) {
	for _, url := range list.GetUrls("berrybase") {
		b.c.Visit(url)
	}

	b.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://www.berrybase.de"),
	}
	// b.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Berry Base: ", r.URL.String())
	// })

	b.c.OnHTML(".product--detail-upper", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.Id = uuid.New()
		item.Name = e.ChildText(".product--title")
		item.InStock = e.ChildText("#buy-button") == "In den Warenkorb"
		item.PriceString = e.ChildText(".price--content")
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

}

func (b *BerryBase) Wait() {
	b.c.Wait()
}
