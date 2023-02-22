package service

import "github.com/gocolly/colly"

func NewCollector() *colly.Collector {

	c := colly.NewCollector(
		colly.MaxDepth(4),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})

	return c
}
