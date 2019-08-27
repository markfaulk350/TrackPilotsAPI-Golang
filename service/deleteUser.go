package service

import "fmt"

func (svc ServiceImpl) DeleteUser(userID string) error {
	if _, err := svc.GetUser(userID); err != nil {
		fmt.Println("Unable to delete user: " + userID + ". No such user exists with that id.")
		return err
	}

	sqlStatement := `DELETE FROM pilots WHERE id=?;`
	if _, err := svc.DBClient.Exec(sqlStatement, userID); err != nil {
		fmt.Println("Unable to delete user:", userID)
		return err
	}
	return nil
}
