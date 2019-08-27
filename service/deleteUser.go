package service

func (svc ServiceImpl) DeleteUser(userID string) error {
	if _, err := svc.GetUser(userID); err != nil {
		svc.Logger.Error().Err(err).Msg("Get user failed")
		return err
	}

	sqlStatement := `DELETE FROM pilots WHERE id=?;`
	if _, err := svc.DBClient.Exec(sqlStatement, userID); err != nil {
		svc.Logger.Error().Err(err).Msg("Delete user failed")
		return err
	}
	return nil
}
