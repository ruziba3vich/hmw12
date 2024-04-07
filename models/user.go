package models

type User struct {
	id       int    `json:"id"`
	username string `json:"username"`
	email    string `json:"email"`
	password string `json:"password"`
}

func (u User) GetId() int {
	return u.id
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetPassword() string {
	return u.password
}


func (u * User) SetId(id int) {
	u.id = id
}

func (u * User) SetUsername(username string) {
	u.username = username
}

func (u * User) SetEmail(email string) {
	u.email = email
}

func (u * User) SetPassword(password string) {
	u.password = password
}
