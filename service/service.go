package service

import (
	"database/sql"

	"github.com/markfaulk350/TrackPilotsAPI/dbclient"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

type Service interface {
	// Users
	CreateUser(user entity.User) (string, error)
	GetUser(userID string) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(userID string, user entity.User) error
	DeleteUser(userID string) error
	// Groups
	CreateGroup(group entity.Group) (string, error)
	GetGroup(groupID string) (entity.Group, error)
	GetAllGroups() ([]entity.Group, error)
	UpdateGroup(groupID string, group entity.Group) error
	DeleteGroup(groupID string) error
	// Roster
	AddToRoster(roster entity.Roster) error
	RemoveFromRoster(roster entity.Roster) error
	GetGroupRoster(groupID string) ([]entity.User, error)
	// Pings
	GetUsersPings(userID int) ([]entity.Ping, error)
	GetGroupTrackingData(groupID string) ([]entity.UserAndPings, error)
	DiscoverNewTrackingData(users []entity.User) error
	MakeApiCall(user entity.User, whenToQueryFrom int64) error
	RetreiveDataFromSpot(user entity.User, whenToQueryFrom int64) error
	RetreiveDataFromGarmin(user entity.User, whenToQueryFrom int64) error
}

type ServiceImpl struct {
	DBClient *sql.DB
}

func New(config *dbclient.Config) Service {
	DatabaseInstance := dbclient.New(config)
	return ServiceImpl{DBClient: DatabaseInstance}
}
