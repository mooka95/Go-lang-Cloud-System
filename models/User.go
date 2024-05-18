package models
type User struct{
	email string
	password string
	firstName string
	lastName string
	identifier string
}
func NewUser(email,password,firstName,lastName string) *User{
	return &User{
		email: email,
		password: password,
		firstName: firstName,
		lastName: lastName,
	}

}
func (user *User)SetIdentifier(identifier string){
	user.identifier =identifier
}