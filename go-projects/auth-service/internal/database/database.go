package database

import (
	"fmt"

	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/login/user"
)

type DatabaseType int64

const (
    Memory DatabaseType = 0
)

type Database interface {
    // Functions for User
    GetUserByUsername(string) (user.User, error)
    SetUser(user.User) (string, error)
    // Functions for session
    GetSessionIDByUsername(string) (string, error)
    GetUsernameBySessionID(string) (string, error)
    SetSession(string, string) (string, error)
}

type InMemoryDatabase struct {
   userMap map[string]user.User
}

func NewDatabase(dt DatabaseType) (Database, error) {
    var db Database

    switch dt {
        case Memory:
            db = &InMemoryDatabase{}
        default:
            return nil, fmt.Errorf("invalid database type passed in")
    }

    return db, nil

}

func (imd *InMemoryDatabase) GetUserByUsername(username string) (user.User, error) {
    gotUser, exists := imd.userMap[username]
    if !exists {
        return nil, fmt.Errorf("could not find user for username %v", username)
    }
    return gotUser, nil
}

func (imd *InMemoryDatabase) SetUser(user.User) (string, error) {
    // to implement
    return "", nil
}

func (imd *InMemoryDatabase) GetSessionIDByUsername(username string) (string, error) {
    // to implement
    return "", nil
}

func (imd *InMemoryDatabase) GetUsernameBySessionID(sessionID string) (string, error) {
    // to implement
    return "", nil
}

func (imd *InMemoryDatabase) SetSession(username, sessionID string) (string, error) {
    // to implement
    return "", nil
}

