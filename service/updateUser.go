package service

import (
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) UpdateUser(userID string, u entity.User) error {
	if _, err := svc.GetUser(userID); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to update user. User does not exist")
		return err
	}

	sqlStatement := `UPDATE pilots SET fName=?, lName=?, email=?, phone=?, country=?, trkLink=?, trkType=?, gldBrand=?, gldMake=?, gldColor=? WHERE id=?;`

	if _, err := svc.DBClient.Exec(sqlStatement, u.Fname, u.Lname, u.Email, u.Phone, u.Country, u.Trklink, u.Trktype, u.GliderBrand, u.GliderMake, u.GliderColor, userID); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to update user")
		return err
	}
	return nil
}
