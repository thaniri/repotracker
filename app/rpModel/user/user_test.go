package user

import (
	"testing"
)

type setNameTestCase struct {
	user     User
	email    string
	expected string
}

var setNameTestCases = []setNameTestCase{
	{
		User{email: "test", passwordHash: "password123"},
		"John",
		"John",
	},
}

func TestSetName(t *testing.T) {
	for _, test := range setNameTestCases {
		test.user.SetEmail(test.email)
		if test.user.email != test.email {
			t.Errorf("fuck")
		}
	}
}

type setPasswordTestCase struct {
	user     User
	email    string
	expected string
}

var setPasswordTestCases = []setPasswordTestCase{
	{
		User{email: "test", passwordHash: "password123"},
		"secureit!",
		"secureit!",
	},
}

func TestSetPasswordHash(t *testing.T) {
	testUser := User{email: "test", passwordHash: "password123"}
	testUser.SetPasswordHash("secure")
	if testUser.passwordHash != "secure" {
		t.Errorf("Username is not secure!!")
	}

}
