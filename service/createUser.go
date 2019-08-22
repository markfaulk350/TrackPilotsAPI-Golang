package service

import (
	"fmt"
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) CreateUser(u entity.User) (string, error) {
	sqlStatement := `INSERT INTO pilots(fName, lName, email, phone, country, trkLink, trkType, gldBrand, gldMake, gldColor) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := svc.DBClient.Exec(sqlStatement, u.Fname, u.Lname, u.Email, u.Phone, u.Country, u.Trklink, u.Trktype, u.GliderBrand, u.GliderMake, u.GliderColor)
	if err != nil {
		fmt.Println("Unable to create user.")
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Unable to grab id of recentley created user.")
		return "", err
	}
	newUserID := strconv.FormatInt(id, 10)
	return newUserID, nil
}
