package service

import (
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetGroupTrackingData(groupID string) ([]entity.UserAndPings, error) {

	users, err := svc.GetGroupRoster(groupID)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to retrieve group roster while getting group tracking data")
		return nil, err
	}

	var usersAndPings []entity.UserAndPings

	for _, user := range users {
		var thisUserWithPings entity.UserAndPings

		thisUserWithPings.User = user

		pingSQLStatement := `SELECT id, unixTime, lat, lng, alt, agl, velocity, heading, txtMsg, isEmergency FROM pings WHERE pilot_id=? ORDER BY unixTime DESC;`
		rows, err := svc.DBClient.Query(pingSQLStatement, user.ID)
		if err != nil {
			svc.Logger.Error().Err(err).Msg("Failed to retrieve user tracking data while retriving group tracking data")
			return nil, err
		}

		for rows.Next() {
			var p entity.Ping
			if err := rows.Scan(&p.ID, &p.UnixTime, &p.Lat, &p.Lng, &p.Alt, &p.AGL, &p.Velocity, &p.Heading, &p.TxtMsg, &p.IsEmergency); err != nil {
				svc.Logger.Error().Err(err).Msg("Failed looping through users tracking data")
				return nil, err
			}
			thisUserWithPings.Pings = append(thisUserWithPings.Pings, p)
		}
		usersAndPings = append(usersAndPings, thisUserWithPings)
	}
	return usersAndPings, nil
}
