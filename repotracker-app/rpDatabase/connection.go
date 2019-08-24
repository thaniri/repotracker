package rpDatabase

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thaniri/repotracker-app/rpLogger"
)

var database Database
var ConnectionString string

// Database struct contains all the required properties to make a connection string for the mysql driver
type Database struct {
	IP       string
	User     string
	Password string
	Database string
	Extra    string
}

func init() {
	readConfig()
}

// Reads the database config from a file and forms a connection string to be used for queries
func readConfig() {
	file, err := ioutil.ReadFile("./config/databaseConfig.json")
	checkErr(err)

	json.Unmarshal([]byte(file), &database)

	ConnectionString = (database.User +
		":" +
		database.Password +
		"@" +
		"tcp(" +
		database.IP +
		")/" +
		database.Database +
		database.Extra)

	rpLogger.Logger.Info("Read database config from file: ", database.Database)
	rpLogger.Logger.Flush()
}

// Takes in an SQL query as a string, and any number of arguments as parameters to that query.
// Has no return
// TODO: return error at least
func ExecuteQuery(query string, args ...interface{}) {
	db, err := sql.Open("mysql", ConnectionString)
	checkErr(err)
	defer db.Close()

	statement, err := db.Prepare(query)
	checkErr(err)

	result, err := statement.Exec(args...)
	checkErr(err)

	// TODO: make this human readable in the logs
	rpLogger.Logger.Info(result)
}

// Takes in an SQL select query as a string
// Example: ExecuteQuery("SELECT ? from rpUsers, "email"
func SelectQuery(query string, args ...interface{}) {
	db, err := sql.Open("mysql", ConnectionString)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query(query, args...)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan()
	}

}

func checkErr(err error) {
	if err != nil {
		rpLogger.Logger.Warn(err)
	}
}
