package scraping

import (
	"github.com/gocolly/colly"
	"gorm.io/datatypes"
	"lunchbunch/internal/models"
	"strconv"
	"strings"
	"time"
)

func Scrape(c *colly.Collector) (result models.RestaurantSnapshot, err error) {
	result.Date = datatypes.Date(time.Now())
	result.Restaurants, err = parseRestaurants(c)
	return
}

func parseRestaurants(c *colly.Collector) (restaurants []models.Restaurant, err error) {
	restaurants = []models.Restaurant{}

	c.OnHTML("div[class='menicka_detail']", func(restaurantElement *colly.HTMLElement) {
		var name = restaurantElement.ChildText("div[class='info'] > div[class='nazev'] > a")

		var result = models.Restaurant{Name: name, MenuItems: parseMenuItems(restaurantElement)}
		restaurants = append(restaurants, result)
	})

	err = c.Visit("https://www.menicka.cz/brno.html")

	return
}

func parseMenuItems(restaurantElement *colly.HTMLElement) (menuItems []models.MenuItem) {
	menuItems = []models.MenuItem{}

	var parseMenuItem = func(iItem int, itemElement *colly.HTMLElement) {
		var itemName = itemElement.Text
		var itemPrice = NoPrice

		if priceElement := itemElement.DOM.NextFiltered("div[class='cena']"); priceElement != nil {
			var priceAndCurrency = strings.SplitN(priceElement.Text(), " ", 2)
			var number, err = strconv.Atoi(priceAndCurrency[0])
			if err == nil {
				itemPrice = number
			}
		}
		var menuitem = models.MenuItem{Name: itemName, Price: itemPrice}
		menuItems = append(menuItems, menuitem)
	}

	restaurantElement.ForEach("div[class='menicka'] > div[class*='nabidka']", parseMenuItem)

	return
}
