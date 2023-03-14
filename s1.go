package s1 

import "fmt"

type User struct {
	firstName string
	lastName  string
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (u *User) SetFirstName(firstname string) {
	u.firstName = firstname
}
func (u *User) SetLastName(lastname string) {
	u.lastName = lastname
}
func (u *User) FullName() string {
	return u.firstName + " " + u.lastName
}

func New() User {
	return User{}
}
func ResetUser(user *User) {
	(*user).firstName = ""
	(*user).lastName = ""
}
func IsUser(user interface{}) bool {
	return fmt.Sprintf("%T", user) == fmt.Sprintf("%T", User{})
}
func ProcessUser(ui UserInterface) string {
	// ui.SetFirstName("ABCDEFGH")
	// ui.SetLastName("JKLMNOP")
	return ui.FullName()
}