package service

import (
	"database/sql"
	"os"

	"github.com/markfaulk350/TrackPilotsAPI/dbclient"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/rs/zerolog"
)

type Service interface {
	// Users
	CreateUser(user entity.User) (entity.CreateUserResult, error)
	GetUser(userID string) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(userID string, user entity.User) error
	DeleteUser(userID string) error
	// Groups
	CreateGroup(group entity.Group) (entity.CreateGroupResult, error)
	GetGroup(groupID string) (entity.Group, error)
	GetAllGroups() ([]entity.Group, error)
	UpdateGroup(groupID string, group entity.Group) error
	DeleteGroup(groupID string) error
	// Roster
	AddToRoster(roster entity.Roster) (entity.CreateRosterResult, error)
	RemoveFromRoster(roster entity.Roster) error
	GetGroupRoster(groupID string) ([]entity.User, error)
	// Pings
	GetUsersPings(userID int) ([]entity.Ping, error)
	GetGroupTrackingData(groupID string, timeSpan string) ([]entity.UserAndPings, error)
}

type ServiceImpl struct {
	DBClient *sql.DB
	Logger   *zerolog.Logger
}

func New(config *dbclient.Config) Service {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	DatabaseInstance := dbclient.New(config)
	return ServiceImpl{
		DBClient: DatabaseInstance,
		Logger:   &logger,
	}
}

type ProfileNotFoundError struct {
	msg string
}

func (e ProfileNotFoundError) Error() string {
	return e.msg
}
