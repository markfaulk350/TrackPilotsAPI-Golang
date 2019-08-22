package service

import (
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) RemoveFromRoster(r entity.Roster) error {
	// Still need to check if user is on group roster before removing them from it.

	sqlStatement := `DELETE FROM groups_have_pilots WHERE group_id=? AND pilot_id=?;`
	_, err := svc.DBClient.Exec(sqlStatement, r.Groupid, r.Pilotid)
	if err != nil {
		fmt.Println("Unable to remove pilot from group roster in DB.")
		return err
	}
	return nil
}
