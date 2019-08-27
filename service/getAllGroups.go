package service

import (
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetAllGroups() ([]entity.Group, error) {
	sqlStatement := `SELECT * FROM flying_groups`
	results, err := svc.DBClient.Query(sqlStatement)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("Failed to retrieve all groups")
		return nil, err
	}

	var allGroups []entity.Group

	for results.Next() {
		var g entity.Group
		err = results.Scan(&g.ID, &g.Groupname, &g.Creatorid, &g.Region, &g.Info, &g.Radio, &g.Created)
		if err != nil {
			svc.Logger.Error().Err(err).Msg("Failed to scan through all groups")
			return nil, err
		}
		allGroups = append(allGroups, g)
	}
	return allGroups, nil
}
