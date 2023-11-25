package scraping

import "github.com/gocolly/colly"

const NoPrice = -1

func ConfigScraper() (c *colly.Collector) {
	c = colly.NewCollector()
	c.AllowURLRevisit = true
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	return
}
