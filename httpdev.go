// Package httpdev implements some http handlers useful for testing
// and especially while developing http servers.
package httpdev

import (
	"fmt"
	"net/http"
	"time"
)

var now = time.Now()

// Handler accepts a "duration" query parameter and will sleep and
// respond with a string message.
func Sleep(w http.ResponseWriter, r *http.Request) {
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	time.Sleep(duration)
	w.Write([]byte(fmt.Sprintf(
		"Started at %s slept for %d nanoseconds.\n",
		now,
		duration.Nanoseconds())))
}
