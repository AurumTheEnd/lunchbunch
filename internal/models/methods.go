package models

import "time"

func (receiver RestaurantSnapshot) TotalVotes() (total uint) {
	total = 0

	for _, restaurant := range receiver.Restaurants {
		total += uint(len(restaurant.Votes))
	}

	return total
}

func (receiver RestaurantSnapshot) TotalVotesString() string {
	if receiver.TotalVotes() == 1 {
		return "vote"
	}

	return "votes"
}

func (receiver RestaurantSnapshot) DateTimeCreated() string {
	return receiver.Datetime.Format(time.DateTime)
}
