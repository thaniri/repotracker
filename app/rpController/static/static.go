package static

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/thaniri/repotracker-app/rpLogger"
	"net/http"
	"path/filepath"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "test",
		Help: "Nohelp",
	})
)

// Handler for GET requests to /
// TODO: make use of caching because / is just a static page.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := filepath.Abs("./rpView/index.html")
	if err != nil {
		rpLogger.Logger.Warn(err)
	}
	counter.Inc()

	http.ServeFile(w, r, file)
}

// Handler for GET requests to /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	file, err := filepath.Abs("./rpView/login.html")
	if err != nil {
		rpLogger.Logger.Warn("test")
	}

	http.ServeFile(w, r, file)
}
