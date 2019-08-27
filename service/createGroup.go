package service

import (
	"strconv"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) CreateGroup(g entity.Group) (entity.CreateGroupResult, error) {
	sqlStatement := `INSERT INTO flying_groups(groupName, creatorId, region, info, radioFrq) VALUES (?, ?, ?, ?, ?)`

	result, err := svc.DBClient.Exec(sqlStatement, g.Groupname, g.Creatorid, g.Region, g.Info, g.Radio)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Create group failed")
		return entity.CreateGroupResult{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed retrieving new group ID")
		return entity.CreateGroupResult{}, err
	}
	newID := strconv.FormatInt(id, 10)
	return entity.CreateGroupResult{GroupID: newID}, nil
}
