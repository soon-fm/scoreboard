// Returns the current build version passed in at compile time
// If no build version was present at compile time then unknown
// will be used as the version value

package version

import "errors"

var (
	version          string // Version string -ldflags "-X scoreboard/version.version=abcdefg"
	ErrVersionNotSet = errors.New("no version set")
)

// Exported method for returning the version string
func Version() (string, error) {
	if version == "" {
		return "", ErrVersionNotSet
	}
	return version, nil
}
