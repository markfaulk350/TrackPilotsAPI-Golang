package service

import (
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) UpdateGroup(groupID string, g entity.Group) error {
	if _, err := svc.GetGroup(groupID); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to update group. Group does not exist")
		return err
	}

	sqlStatement := `UPDATE flying_groups SET groupName=?, creatorId=?, region=?, info=?, radioFrq=? WHERE id=?;`

	if _, err := svc.DBClient.Exec(sqlStatement, g.Groupname, g.Creatorid, g.Region, g.Info, g.Radio, groupID); err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to update group")
		return err
	}
	return nil
}
