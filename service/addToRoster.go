package service

import (
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) AddToRoster(r entity.Roster) (entity.CreateRosterResult, error) {
	// Check if pilot exists
	pilotIDString := strconv.Itoa(r.Pilotid)
	if _, err := svc.GetUser(pilotIDString); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to add user to group. User does not exist")
		return entity.CreateRosterResult{}, err
	}

	// Check if group exists
	groupIDString := strconv.Itoa(r.Groupid)
	if _, err := svc.GetGroup(groupIDString); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to add user to group. Group does not exist")
		return entity.CreateRosterResult{}, err
	}

	// Add user to group
	sqlStatement := `INSERT INTO groups_have_pilots(group_id, pilot_id) VALUES (?, ?);`
	result, err := svc.DBClient.Exec(sqlStatement, r.Groupid, r.Pilotid)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to add user to group roster")
		return entity.CreateRosterResult{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed retrieving new roster ID")
		return entity.CreateRosterResult{}, err
	}
	newID := strconv.FormatInt(id, 10)
	return entity.CreateRosterResult{RosterID: newID}, nil
}
