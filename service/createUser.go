package service

import (
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) CreateUser(u entity.User) (entity.CreateUserResult, error) {
	sqlStatement := `INSERT INTO pilots(fName, lName, email, phone, country, trkLink, trkType, gldBrand, gldMake, gldColor) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := svc.DBClient.Exec(sqlStatement, u.Fname, u.Lname, u.Email, u.Phone, u.Country, u.Trklink, u.Trktype, u.GliderBrand, u.GliderMake, u.GliderColor)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Create user failed")
		return entity.CreateUserResult{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed retrieving new user ID")
		return entity.CreateUserResult{}, err
	}
	newID := strconv.FormatInt(id, 10)
	return entity.CreateUserResult{UserID: newID}, nil
}
