package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
    SessionID string
    Expiry time.Time
}

func NewSession() (*Session, error) {
    id, err := uuid.NewRandom()
    if err != nil {
        return nil, err
    }

    expiration := time.Now().Add(8 * time.Hour)

    session := &Session {
        SessionID: id.String(),
        Expiry: expiration,
    }
    
    // make the database call to store this sessionID

    return session, nil
}

func VerifySession() {
    // yet to implement
}

func DeleteSession() {
    // yet to implement
}
