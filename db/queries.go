package db

import (
	"fmt"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

// Returns weekday (mon-fri) time range for the given date. If the given
// date is a Satuday or Sunday this date is adjusted to return the previous
// weeks time range
func weekdayTimeRange(date time.Time) (mon time.Time, fri time.Time) {
	wd := date.Weekday()
	// If we are not a mon-fri date, adjust date to the previous week
	if wd == time.Saturday {
		date = date.AddDate(0, 0, -1) // friday
		wd = date.Weekday()
	}
	if wd == time.Sunday {
		date = date.AddDate(0, 0, -2) // friday
		wd = date.Weekday()
	}
	mon = date
	for mon.Weekday() != time.Monday {
		mon = mon.AddDate(0, 0, -1)
	}
	fri = date
	for fri.Weekday() != time.Friday {
		fri = fri.AddDate(0, 0, 1)
	}
	return mon, fri.AddDate(0, 0, 1)
}

var scoresByWeekQry = `SELECT SUM("value") AS total
FROM "scores"
WHERE time >= '%s'
	AND time <= '%s'
GROUP BY "user";`

func ScoresByWeek(q Queryer) ([]influxdb.Result, error) {
	now := time.Now().UTC()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	mon, fri := weekdayTimeRange(date)
	qry := fmt.Sprintf(scoresByWeekQry, mon.Format(time.RFC3339), fri.Format(time.RFC3339))
	return q.Query(qry)
}
