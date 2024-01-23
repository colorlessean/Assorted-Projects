package user

import "errors"

type UserID string

// UserInfo is the user information
type UserInfo struct {
	Username UserID `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

// New creates a new user
func New(username UserID, email, bio string) (*UserInfo, error) {
	if username == "" || email == "" {
		return nil, errors.New("cannot create a user without a username or email")
	}

	return &UserInfo{
		Username: username,
		Email:    email,
		Bio:      bio,
	}, nil
}
