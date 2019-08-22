package service

import (
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

// Will grab all users in a specific group and all thier tracking data
func (svc ServiceImpl) GetGroupTrackingData(groupID string) ([]entity.UserAndPings, error) {
	// First need to check if group exists
	_, err := svc.GetGroup(groupID)
	if err != nil {
		fmt.Println("Unable to get groups users and thier location data. Group with ID:" + groupID + " does not exist.")
		return nil, err
	}

	// Then need to get all users in group and their personal data
	sqlStatement := `SELECT pilots.id, fName, lName, email, phone, country, trkLink, trkType, gldBrand, gldMake, gldColor, lastLocationPing, lastApiCall, pilots.created FROM groups_have_pilots INNER JOIN pilots ON groups_have_pilots.pilot_id = pilots.id WHERE groups_have_pilots.group_id=?;`
	rows, err := svc.DBClient.Query(sqlStatement, groupID)
	if err != nil {
		fmt.Println("Unable to grab pilots data from a group tracking query.")
		return nil, err
	}

	var users []entity.User

	for rows.Next() {
		var u entity.User

		err = rows.Scan(&u.ID, &u.Fname, &u.Lname, &u.Email, &u.Phone, &u.Country, &u.Trklink, &u.Trktype, &u.GliderBrand, &u.GliderMake, &u.GliderColor, &u.LastLocationPing, &u.LastApiCall, &u.Created)
		if err != nil {
			fmt.Println("Unable to grab single pilots personal data from a group query.")
			return nil, err
		}
		users = append(users, u)
	}

	// Now we can pass the user with empty pings to a function that will determine who has new tracking data and who needs to be updated
	err = svc.DiscoverNewTrackingData(users)
	if err != nil {
		fmt.Println("Unable to get latest tracking data on group" + groupID)
		return nil, err
	}

	var usersAndPings []entity.UserAndPings

	// After we determine all users are up to date, we can loop through our users and query for tracking data and append it to the userAndPings Struct.
	for _, user := range users {
		var thisUserWithPings entity.UserAndPings

		thisUserWithPings.User = user

		pingSQLStatement := `SELECT id, unixTime, lat, lng, alt, agl, velocity, heading, txtMsg, isEmergency FROM pings WHERE pilot_id=? ORDER BY unixTime DESC;`
		rows, err := svc.DBClient.Query(pingSQLStatement, user.ID)
		if err != nil {
			fmt.Println("Unable to grab single pilots location data from a group query.")
			return nil, err
		}

		for rows.Next() {
			var p entity.Ping
			err := rows.Scan(&p.ID, &p.UnixTime, &p.Lat, &p.Lng, &p.Alt, &p.AGL, &p.Velocity, &p.Heading, &p.TxtMsg, &p.IsEmergency)
			if err != nil {
				fmt.Println("Unable to scan row of location pings.")
				return nil, err
			}
			thisUserWithPings.Pings = append(thisUserWithPings.Pings, p)
		}
		usersAndPings = append(usersAndPings, thisUserWithPings)
	}
	return usersAndPings, err
}
