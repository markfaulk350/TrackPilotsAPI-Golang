package service

import "fmt"

func (svc ServiceImpl) DeleteGroup(groupID string) error {
	_, err := svc.GetGroup(groupID)
	if err != nil {
		fmt.Println("Unable to delete group: " + groupID + ". No such group exists with that id.")
		return err
	}

	sqlStatement := `DELETE FROM flying_groups WHERE id=?;`
	_, err = svc.DBClient.Exec(sqlStatement, groupID)
	if err != nil {
		fmt.Println("Unable to delete group:", groupID)
		return err
	}
	return nil
}
