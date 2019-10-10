package service

import (
	"errors"
	"time"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetGroupTrackingData(groupID string, timeSpan string) ([]entity.UserAndPings, error) {

	currentUnix := time.Now().Unix()

	var pingSQLStatement = `SELECT id, unixTime, lat, lng, alt, agl, velocity, heading, txtMsg, isEmergency FROM pings WHERE pilot_id=? AND unixTime >=? ORDER BY unixTime DESC;`
	var queryFrom int64
	var queryForMostRecent bool

	switch timeSpan {
	case "mostrecent":
		pingSQLStatement = `SELECT id, MAX(unixTime), lat, lng, alt, agl, velocity, heading, txtMsg, isEmergency FROM pings WHERE pilot_id=?;`
		queryForMostRecent = true
	case "1hr":
		queryFrom = currentUnix - 3600
		queryForMostRecent = false
	case "12hr":
		queryFrom = currentUnix - 43200
		queryForMostRecent = false
	case "48hr":
		queryFrom = currentUnix - 172800
		queryForMostRecent = false
	case "1week":
		queryFrom = currentUnix - 604800
		queryForMostRecent = false
	default:
		return nil, errors.New("time span has not been specified")
	}

	users, err := svc.GetGroupRoster(groupID)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to retrieve group roster while getting group tracking data")
		return nil, err
	}

	var usersAndPings []entity.UserAndPings

	if queryForMostRecent == false {
		for _, user := range users {
			var thisUserWithPings entity.UserAndPings

			thisUserWithPings.User = user

			rows, err := svc.DBClient.Query(pingSQLStatement, user.ID, queryFrom)
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
	} else {
		for _, user := range users {
			var thisUserWithPings entity.UserAndPings

			thisUserWithPings.User = user

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

	}

	return usersAndPings, nil
}

// SELECT id, unixTime, lat, lng, alt, agl, velocity, heading, txtMsg, isEmergency FROM pings WHERE pilot_id=2 AND unixTime >=1569099480 ORDER BY unixTime DESC;
