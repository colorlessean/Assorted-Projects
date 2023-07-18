package user

import "errors"

type User struct {
	username string
	email    string
}

func (u *User) Init(username, email string) error {
	if username == "" || email == "" {
		return errors.New("invalid inputs")
	}

	u.username = username
	u.email = email
	return nil
}

func GetUserByUsername(username string) (User, error) {
	user := User{}

	// call the db call get user by key

	return user, nil
}

func UpdateOrCreateUser(user *User) {

}

func GetAllLikedPosts()

func GetAllPosts()

func MakePost()

func MakeRepost()
