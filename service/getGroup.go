package service

import (
	"database/sql"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetGroup(groupID string) (entity.Group, error) {
	sqlStatement := `SELECT * FROM flying_groups WHERE id=?;`
	row := svc.DBClient.QueryRow(sqlStatement, groupID)

	var g entity.Group

	switch err := row.Scan(&g.ID, &g.Groupname, &g.Creatorid, &g.Region, &g.Info, &g.Radio, &g.Created); err {
	case sql.ErrNoRows:
		svc.Logger.Error().Err(err).Msg("Failed to retrieve group. Could not find group with ID: " + groupID)
		return entity.Group{}, ProfileNotFoundError{"Could not find group with ID: " + groupID}
	case nil:
		return g, nil
	default:
		svc.Logger.Error().Err(err).Msg("Failed to retrieve group")
		return entity.Group{}, err
	}
}
