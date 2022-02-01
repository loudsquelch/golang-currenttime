package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05 -0700 MST"

type currentTimeDto struct {
	Time string `json:"current_time"`
}

func now(w http.ResponseWriter, r *http.Request) {

	// Retrieve URL parameter
	tzq := r.URL.Query().Get("tz")

	times := make(map[string]string)
	for _, tz := range strings.Split(tzq, ",") {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.Header().Add("Content-Type", "text/plain")
			// Prefer 400 instead of 404
			http.Error(w, "invalid timezone", http.StatusBadRequest)
			return
		}
		times[tz] = time.Now().In(loc).Format(timeFormat)
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	if strings.ContainsAny(tzq, ",") {
		enc.Encode(times)
	} else {
		enc.Encode(currentTimeDto{Time: times[tzq]})
	}

}
