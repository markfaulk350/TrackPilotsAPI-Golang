package service

import (
	"database/sql"
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

// Eventually this should take an interface so it can take strings and ints
func (svc ServiceImpl) GetUser(userID string) (entity.User, error) {

	sqlStatement := `SELECT * FROM pilots WHERE id=?;`
	row := svc.DBClient.QueryRow(sqlStatement, userID)

	var u entity.User

	switch err := row.Scan(&u.ID, &u.Fname, &u.Lname, &u.Email, &u.Phone, &u.Country, &u.Trklink, &u.Trktype, &u.GliderBrand, &u.GliderMake, &u.GliderColor, &u.LastLocationPing, &u.LastApiCall, &u.Created); err {
	case sql.ErrNoRows:
		fmt.Println("Could not find user with ID of", userID)
		return entity.User{}, err
	case nil:
		fmt.Println(u)
		return u, nil
	default:
		fmt.Println("Unable to grab users data. Something went wrong.")
		return entity.User{}, err
	}
}
