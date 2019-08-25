package user

type User struct {
	email        string
	passwordHash string
}

// Setter method for user email
func (user *User) SetEmail(email string) {
	user.email = email
}

// Setter method for user passwordHash
func (user *User) SetPasswordHash(passwordHash string) {
	user.passwordHash = passwordHash
}
