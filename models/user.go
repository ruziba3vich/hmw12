package models

type User struct {
	id       int   
	username string
	email    string
	password string
	age      int   
}


/*
Pastda user ni saqlash uchun unga getter va setterlar yaratildi
*/

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

func (u User) GetAge() int {
	return u.age
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetAge(age int) {
	u.age = age
}
