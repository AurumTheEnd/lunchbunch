package models

func (receiver RestaurantSnapshot) TotalVotes() (total uint) {
	total = 0

	for _, restaurant := range receiver.Restaurants {
		total += restaurant.Votes
	}

	return
}

func (receiver RestaurantSnapshot) TotalVotesString() string {
	if receiver.TotalVotes() == 1 {
		return "vote"
	}

	return "votes"
}
