package models

import "time"

type VoteCount interface {
	TotalVotes() (total uint)
	TotalVotesString() string
}

func (receiver RestaurantSnapshot) TotalVotes() (total uint) {
	total = 0

	for _, restaurant := range receiver.Restaurants {
		total += restaurant.TotalVotes()
	}

	return total
}

func (receiver Restaurant) TotalVotes() (total uint) {
	return uint(len(receiver.Votes))
}

func (receiver RestaurantSnapshot) TotalVotesString() string {
	return totalVotesString(receiver.TotalVotes())
}

func (receiver Restaurant) TotalVotesString() string {
	return totalVotesString(receiver.TotalVotes())
}

func totalVotesString(count uint) string {
	if count == 1 {
		return "vote"
	}

	return "votes"
}

func (receiver RestaurantSnapshot) PercentageOfVotes(restaurant Restaurant) uint {
	return restaurant.TotalVotes() / receiver.TotalVotes() * 100
}

func (receiver RestaurantSnapshot) DateTimeCreated() string {
	return receiver.Datetime.Format(time.DateTime)
}

func (receiver Restaurant) HasUserVoted(userID uint) bool {
	for _, vote := range receiver.Votes {
		if vote.UserID == userID {
			return true
		}
	}

	return false
}
