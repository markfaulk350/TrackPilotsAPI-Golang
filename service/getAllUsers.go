package service

import (
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetAllUsers() ([]entity.User, error) {
	sqlStatement := `SELECT * FROM pilots`
	results, err := svc.DBClient.Query(sqlStatement)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to retrieve all users")
		return nil, err
	}

	var allUsers []entity.User

	for results.Next() {
		var u entity.User
		err = results.Scan(&u.ID, &u.Fname, &u.Lname, &u.Email, &u.Phone, &u.Country, &u.Trklink, &u.Trktype, &u.GliderBrand, &u.GliderMake, &u.GliderColor, &u.LastLocationPing, &u.LastApiCall, &u.Created)
		if err != nil {
			svc.Logger.Error().Err(err).Msg("Failed to scan through all users")
			return nil, err
		}
		allUsers = append(allUsers, u)
	}
	return allUsers, nil
}
