package service

import (
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) UpdateGroup(groupID string, g entity.Group) error {
	_, err := svc.GetGroup(groupID)
	if err != nil {
		fmt.Println("Unable to update group. Group: " + groupID + " does not exist.")
		return err
	}
	sqlStatement := `UPDATE flying_groups SET groupName=?, creatorId=?, region=?, info=?, radioFrq=? WHERE id=?;`
	_, err = svc.DBClient.Exec(sqlStatement, g.Groupname, g.Creatorid, g.Region, g.Info, g.Radio, groupID)
	if err != nil {
		fmt.Println("Unable to update group info.")
		return err
	}
	return nil
}
