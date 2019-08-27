package service

import (
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetGroupRoster(groupID string) ([]entity.User, error) {
	if _, err := svc.GetGroup(groupID); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to retrieve group roster. Group does not exist")
		return nil, err
	}

	sqlStatement := `SELECT pilots.id, fName, lName, email, phone, country, trkLink, trkType, gldBrand, gldMake, gldColor, lastLocationPing, lastApiCall, pilots.created FROM groups_have_pilots INNER JOIN pilots ON groups_have_pilots.pilot_id = pilots.id WHERE groups_have_pilots.group_id=?;`
	rows, err := svc.DBClient.Query(sqlStatement, groupID)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to retrieve group roster")
		return nil, err
	}

	var groupRoster []entity.User

	for rows.Next() {
		var u entity.User
		err = rows.Scan(&u.ID, &u.Fname, &u.Lname, &u.Email, &u.Phone, &u.Country, &u.Trklink, &u.Trktype, &u.GliderBrand, &u.GliderMake, &u.GliderColor, &u.LastLocationPing, &u.LastApiCall, &u.Created)
		if err != nil {
			svc.Logger.Error().Err(err).Msg("Failed to scan through group roster")
			return nil, err
		}
		groupRoster = append(groupRoster, u)
	}
	return groupRoster, nil
}
