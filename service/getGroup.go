package service

import (
	"database/sql"
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetGroup(groupID string) (entity.Group, error) {
	sqlStatement := `SELECT * FROM flying_groups WHERE id=?;`
	row := svc.DBClient.QueryRow(sqlStatement, groupID)

	var g entity.Group

	switch err := row.Scan(&g.ID, &g.Groupname, &g.Creatorid, &g.Region, &g.Info, &g.Radio, &g.Created); err {
	case sql.ErrNoRows:
		fmt.Println("Could not find group with ID of", groupID)
		return entity.Group{}, ProfileNotFoundError{"Could not find group with ID: " + groupID}
	case nil:
		//fmt.Println(g)
		return g, nil
	default:
		fmt.Println("Unable to grab groups data.")
		return entity.Group{}, err
	}
}
