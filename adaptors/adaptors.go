package adaptors

import "github.com/jufabeck2202/piScraper/utils"

type Adaptor interface {
	Run(list utils.Websites)
	Wait()
}
