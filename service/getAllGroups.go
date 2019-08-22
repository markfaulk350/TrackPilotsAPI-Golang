package service

import (
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) GetAllGroups() ([]entity.Group, error) {
	sqlStatement := `SELECT * FROM flying_groups`
	results, err := svc.DBClient.Query(sqlStatement)
	if err != nil {
		fmt.Println("Unable to grab all Groups")
		return nil, err
	}

	var allGroups []entity.Group

	for results.Next() {
		var g entity.Group
		err = results.Scan(&g.ID, &g.Groupname, &g.Creatorid, &g.Region, &g.Info, &g.Radio, &g.Created)
		if err != nil {
			fmt.Println("Unable scan all Groups data from DB")
			return nil, err
		}
		allGroups = append(allGroups, g)
	}
	return allGroups, nil
}
