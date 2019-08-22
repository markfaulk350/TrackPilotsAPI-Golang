package service

import (
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) UpdateUser(userID string, u entity.User) error {
	_, err := svc.GetUser(userID)
	if err != nil {
		fmt.Println("Unable to update user. User: " + userID + " does not exist.")
		return err
	}

	sqlStatement := `UPDATE pilots SET fName=?, lName=?, email=?, phone=?, country=?, trkLink=?, trkType=?, gldBrand=?, gldMake=?, gldColor=? WHERE id=?;`
	_, err = svc.DBClient.Exec(sqlStatement, u.Fname, u.Lname, u.Email, u.Phone, u.Country, u.Trklink, u.Trktype, u.GliderBrand, u.GliderMake, u.GliderColor, userID)
	if err != nil {
		fmt.Println("Unable to update user info.")
		return err
	}
	return nil
}
