package user

type User struct {
	email        string
	passwordHash string
}

// Setter method for user email
func (user *User) SetName(email string) {
	user.email = email
}

// Getter method for user email
func (user User) Email() string {
	return user.email
}

// Setter method for user passwordHash
func (user *User) SetPasswordHash(passwordHash string) {
	user.passwordHash = passwordHash
}

// Getter method for user passwordHash
func (user User) PasswordHash() string {
	return user.passwordHash
}
