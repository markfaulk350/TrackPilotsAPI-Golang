package service

import (
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) RemoveFromRoster(r entity.Roster) error {
	// Check if pilot exists
	pilotIDString := strconv.Itoa(r.Pilotid)
	if _, err := svc.GetUser(pilotIDString); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to remove user from group. User does not exist")
		return err
	}

	// Check if group exists
	groupIDString := strconv.Itoa(r.Groupid)
	if _, err := svc.GetGroup(groupIDString); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to remove user from group. Group does not exist")
		return err
	}

	sqlStatement := `DELETE FROM groups_have_pilots WHERE group_id=? AND pilot_id=?;`
	if _, err := svc.DBClient.Exec(sqlStatement, r.Groupid, r.Pilotid); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to remove user from group")
		return err
	}
	return nil
}
