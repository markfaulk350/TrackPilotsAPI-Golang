package service

import (
	"fmt"
	"time"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/helpers"
)

// This function takes in a struct of users and determines if we already have the most up to date tracking information on them.
// If we have all the latest tracking info then there is no need to make any API requests to SPOT or Garmin Inreach.
// If there is a possibility we do not have the most up to date information on a users location, then we will have to make a request to the users tracking url.
// Once we get the users new tracking data we will have to parse it, normalize the format, then save it to our database for future use.

// The main purpose of this function is to reduce the amount repeat, unnecessary API calls to both SPOT and Garmin Inreach.
// If we make too many API calls to either of the two services, we may get blocked.

// Loop through each user, look at lastAPICall and LastLocationPing to determine if we need to make a request for new tracking data
func (svc ServiceImpl) DiscoverNewTrackingData(users []entity.User) error {
	currentUnixTime := time.Now().Unix()
	for _, user := range users {
		secondsSinceLastApiCall := currentUnixTime - user.LastApiCall
		secondsSinceLastLocationPing := currentUnixTime - user.LastLocationPing

		var whenToQueryFrom int64

		// Determine how much data we already have from the last api call and query from then to now
		switch {
		case secondsSinceLastApiCall > 0 && secondsSinceLastApiCall < 604800:
			whenToQueryFrom = user.LastApiCall
		default:
			whenToQueryFrom = currentUnixTime - 604800
		}

		// Determine which pilots are already up to date, which are in timeout, and which need to be updated
		switch {
		// If we have a new pilot, then we need to get all their data
		case user.LastApiCall == 0:
			fmt.Println("This is a new user that needs to be updated!")
			svc.MakeApiCall(user, whenToQueryFrom)

		// If a pilot has not been queried for 15 min, it doesnt matter if they are active or not, they need to be checked on
		case secondsSinceLastApiCall > 900:
			fmt.Println("This pilot has not been updated for at least 15 min and will be updated")
			svc.MakeApiCall(user, whenToQueryFrom)

		// If we just checked on a pilot less than 15 min ago and they were not active within the last 30 min or they have never been active, then we can assume they are inactive pilots who do not need to be updated
		case secondsSinceLastApiCall < 900 && (secondsSinceLastLocationPing > 1800 || user.LastLocationPing == 0):
			fmt.Println("This is an inactive pilot who was just updated and is in 15min timeout and will not be updated")

		// We need a case in which the pilot is active, with no new data expected, lastApiCall could be anytime, lastLocationPing would be less than 10 min ago
		case secondsSinceLastLocationPing < 600:
			fmt.Println("Active pilot who will not have any new data avaliable due to 10 min tracking interval")

		// If we recentley checked on a pilot who was expected to have new data less than 10 seconds ago, then there is no need to check again so soon.
		case secondsSinceLastApiCall < 10 && 600 < secondsSinceLastLocationPing && secondsSinceLastLocationPing < 1800:
			fmt.Println("Active pilot with new data expected but not received in 10 second timeout")

		// If a pilot has not sent a ping in the last 10 min and is considered to be active due to a recent location ping then we know that they need to be updated
		case (10 < secondsSinceLastApiCall && secondsSinceLastApiCall < 900) && (600 < secondsSinceLastLocationPing && secondsSinceLastLocationPing < 1800):
			fmt.Println("Active pilot with new location data available that is not in timeout")
			svc.MakeApiCall(user, whenToQueryFrom)
		}
		// Also need a case for a pilot with a tracking link that does not work!

		fmt.Println(whenToQueryFrom)
		fmt.Println(user.Fname + "'s last location ping was " + helpers.Int64ToString(secondsSinceLastLocationPing) + " sec ago")
		fmt.Println("We checked " + user.Fname + "'s location data " + helpers.Int64ToString(secondsSinceLastApiCall) + " sec ago")
	}

	return nil
}
