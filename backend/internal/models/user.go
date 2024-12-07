package models

type User struct {
	Username string
	Password string
}

var users = []User{
	{Username: "alice", Password: "password123"},
	{Username: "bob", Password: "securepassword"},
}

var GetUser(username, password string) *User {
	for _, user := range user {
		if user.Username == username && user.Password == password {
			return &user
		}
	}
	return nil
}