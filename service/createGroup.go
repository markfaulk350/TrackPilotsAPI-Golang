package service

import (
	"fmt"
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) CreateGroup(g entity.Group) (string, error) {
	sqlStatement := `INSERT INTO flying_groups(groupName, creatorId, region, info, radioFrq) VALUES (?, ?, ?, ?, ?)`

	result, err := svc.DBClient.Exec(sqlStatement, g.Groupname, g.Creatorid, g.Region, g.Info, g.Radio)
	if err != nil {
		fmt.Println("Unable to create group.")
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
