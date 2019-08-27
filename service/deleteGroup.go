package service

func (svc ServiceImpl) DeleteGroup(groupID string) error {
	if _, err := svc.GetGroup(groupID); err != nil {
		svc.Logger.Error().Err(err).Msg("Get group failed")
		return err
	}

	sqlStatement := `DELETE FROM flying_groups WHERE id=?;`
	if _, err := svc.DBClient.Exec(sqlStatement, groupID); err != nil {
		svc.Logger.Error().Err(err).Msg("Delete group failed")
		return err
	}
	return nil
}
