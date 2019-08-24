package login

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/thaniri/repotracker-app/rpDatabase"
	"github.com/thaniri/repotracker-app/rpLogger"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Generate new keys for cookies at runtime
// Cookies will be invalidated at application restart
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// Temp until I figure out an intelligent way to do the view
const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

// Creates a basic cookie of a user email
func setSession(email string, writer http.ResponseWriter) {
	value := map[string]string{
		"email": email,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:    "session",
			Value:   encoded,
			Path:    "/",
			Expires: time.Now().Add(4 * time.Hour),
		}
		http.SetCookie(writer, cookie)
	}
}

// Sets the session cookie to expiry date to the past which invalidates it.
func clearSession(writer http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(writer, cookie)
}

// Looks at the cookies for the user making a request. If the session cookie exists, and hasn't expired
// this function will return the value for the email property of the cookie.
// TODO: test cookie manipulation
func getEmail(request *http.Request) (email string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			email = cookieValue["email"]
		}
	}
	return email
}

// Example of an internal page that can only be accessed if a session cookies is properly set.
func InternalPageHandler(writer http.ResponseWriter, request *http.Request) {
	email := getEmail(request)
	if email != "" {
		fmt.Fprintf(writer, internalPage, email)
	} else {
		http.Redirect(writer, request, "/", 302)
	}
}

// Handler for POST requests to /login based on the login form.
func LoginPostHandler(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	redirectTarget := "/"
	if email != "" && password != "" {
		rpLogger.Logger.Debug(email, " login attempt.")
		if checkPasswordHash(email, password) {
			rpLogger.Logger.Info(email, " login succeeded.")
			setSession(email, writer)
			redirectTarget = "/internal"
		} else {
			rpLogger.Logger.Warn(email, " login failed.")
			http.Error(writer, "Login credentials incorrect", 400)
		}

	}
	http.Redirect(writer, request, redirectTarget, 302)
}

// Handler for GET requests to /logout, which clears user cookies.
// TODO: log which user actually logged out
func LogoutPostHandler(writer http.ResponseWriter, request *http.Request) {
	clearSession(writer)
	http.Redirect(writer, request, "/", 302)
}

// Handler for POST requests to /register based on the register HTML form.
// TODO: graceful error handling if for some reason this fails (database unavailable or something)
func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := hashPassword(request.FormValue("password"))
	redirectTarget := "/"
	if email != "" && password != "" {
		query := "INSERT rpUsers SET email=?, passwordHash=?"
		rpDatabase.ExecuteQuery(query, email, password)
		rpLogger.Logger.Info("Registered user")
	}
	http.Redirect(writer, request, redirectTarget, 302)
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		rpLogger.Logger.Error("Failed to generate a hash", err)
	}
	return string(hash)
}

func checkPasswordHash(email string, password string) bool {
	query := "SELECT email, passwordHash from rpUsers WHERE email=?"
	rpDatabase.ExecuteQuery(query, email)
	return true
}
