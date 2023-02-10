package dateExtractor

import (
	"time"
)

func IsLinkExpired(dateTime time.Time) bool {
	currentTime := time.Now()
	return currentTime.Sub(dateTime).Hours() >= 30*24
}
