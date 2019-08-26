package service

import (
	"fmt"
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) AddToRoster(r entity.Roster) (string, error) {
	// Check if pilot exists
	pilotIDString := strconv.Itoa(r.Pilotid)
	_, err := svc.GetUser(pilotIDString)
	if err != nil {
		fmt.Println("Unable to add user to group. User does not exist.")
		return "", err
	}

	// Check if group exists
	groupIDString := strconv.Itoa(r.Groupid)
	_, err = svc.GetGroup(groupIDString)
	if err != nil {
		fmt.Println("Unable to add user to group. Group does not exist.")
		return "", err
	}

	// Add user to group
	sqlStatement := `INSERT INTO groups_have_pilots(group_id, pilot_id) VALUES (?, ?);`
	result, err := svc.DBClient.Exec(sqlStatement, r.Groupid, r.Pilotid)
	if err != nil {
		fmt.Println("Unable to add pilot to group inside DB.")
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Unable to grab id of recentley created group.")
		return "", err
	}
	newID := strconv.FormatInt(id, 10)
	return newID, nil
}