package version

import (
	"errors"
	"strconv"
	"time"
)

var (
	buildTime         string                            // Unix epoch - -ldflags "-X scoreboard/buildTime.buildTime=1482510310"
	ErrNoBuildTimeSet = errors.New("no build time set") // Error if build time not set
)

// Exported method for returning the build time
func BuildTime() (time.Time, error) {
	if buildTime == "" {
		return time.Time{}, ErrNoBuildTimeSet
	}
	i, err := strconv.ParseInt(buildTime, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	t := time.Unix(i, 0)
	return t, nil
}
