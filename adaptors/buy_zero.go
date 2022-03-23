package adaptors

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/jufabeck2202/piScraper/utils"
)

type BuyZero struct {
	c *colly.Collector
}

func NewBuyZero(c *colly.Collector) *BuyZero {
	copy := c.Clone()
	return &BuyZero{copy}
}

func (s *BuyZero) Run(list utils.Websites) {
	for _, url := range list.GetUrls("buyzero") {
		s.c.Visit(url)
	}
	s.c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile("^https://buyzero.de/"),
	}
	// s.c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting Buyzero: ", r.URL.String())
	// })

	s.c.OnHTML(".product-meta__title.heading.h1", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.Name = e.Text
		item.Time = time.Now().Format("15:04:05")
		list.UpdateItemInList(item)
	})

	s.c.OnHTML(".price-list", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.PriceString = strings.Replace(e.ChildText(".price"), "Angebotspreis", "", -1)
		list.UpdateItemInList(item)
	})

	s.c.OnHTML(".product-form__inventory.inventory", func(e *colly.HTMLElement) {
		item := list.GetItemById(e.Request.URL.String())
		item.InStock = e.Text != "Ausverkauft"
		list.UpdateItemInList(item)
	})
}

func (s *BuyZero) Wait() {
	s.c.Wait()
}
