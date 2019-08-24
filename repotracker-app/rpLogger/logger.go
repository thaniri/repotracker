package rpLogger

import (
	"encoding/json"
	"io/ioutil"

	"github.com/chasex/glog"
)

var Logger *glog.Logger

func init() {
	readConfig()
}

// Reads the logger config from a file an makes a Logger object for use globally.
func readConfig() {
	var logConfig glog.LogOptions

	configFile, err := ioutil.ReadFile("./config/loggerConfig.json")

	json.Unmarshal([]byte(configFile), &logConfig)

	Logger, err = glog.New(logConfig)

	if err != nil {
		panic(err)
	}

	Logger.Info("Read logger config from file: ", logConfig)
	Logger.Flush()
}
