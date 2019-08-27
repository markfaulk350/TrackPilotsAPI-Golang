package service

import (
	"errors"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) MakeApiCall(user entity.User, whenToQueryFrom int64) error {
	switch user.Trktype {
	case "spot":
		svc.RetreiveDataFromSpot(user, whenToQueryFrom)
	case "inreach":
		svc.RetreiveDataFromGarmin(user, whenToQueryFrom)
	default:
		svc.Logger.Error().Str("Incorrrect Tracker Type:", user.Trktype).Msg("Tracker is not of type Inreach or SPOT")
		return errors.New("Tracker is not of type Inreach or SPOT")
	}
	return nil
}
