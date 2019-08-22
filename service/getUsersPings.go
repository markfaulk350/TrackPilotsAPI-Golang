package service

import (
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetUsersPings(userID int) ([]entity.Ping, error) {
	pingSQLStatement := `SELECT id, unixTime, lat, lng, alt, agl, velocity, heading, txtMsg, isEmergency FROM pings WHERE pilot_id=? ORDER BY unixTime DESC;`
	rows, err := svc.DBClient.Query(pingSQLStatement, userID)
	if err != nil {
		fmt.Println("Unable to grab single pilots location data.")
		return []entity.Ping{}, err
	}

	var allPings []entity.Ping

	for rows.Next() {
		var p entity.Ping
		err = rows.Scan(&p.ID, &p.UnixTime, &p.Lat, &p.Lng, &p.Alt, &p.AGL, &p.Velocity, &p.Heading, &p.TxtMsg, &p.IsEmergency)
		if err != nil {
			fmt.Println("Unable to scan row of location pings.", err)
		}
		allPings = append(allPings, p)
	}
	return allPings, nil
}
