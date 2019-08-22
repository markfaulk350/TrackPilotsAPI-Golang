package service

import (
	"errors"
	"fmt"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) MakeApiCall(user entity.User, whenToQueryFrom int64) error {
	switch user.Trktype {
	case "spot":
		//fmt.Println("Make API call to SPOT")
		svc.RetreiveDataFromSpot(user, whenToQueryFrom)
	case "inreach":
		//fmt.Println("Make API call to Garmin Inreach")
		svc.RetreiveDataFromGarmin(user, whenToQueryFrom)
	default:
		fmt.Println("Problem making API call. Tracker type doesnt seem to be of type Garmin Inreach or SPOT.")
		return errors.New("problem making API call. Tracker type doesnt seem to match Garmin Inreach or SPOT")
	}
	return nil
}
