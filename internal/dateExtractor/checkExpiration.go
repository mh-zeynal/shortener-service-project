package dateExtractor

import (
	"strconv"
	"strings"
	"time"
)

func IsLinkExpired(dateTime string) bool {
	space      := strings.Index(dateTime, " ")
	plus       := strings.Index(dateTime, "+")
	year, _    := strconv.Atoi(strings.Split(dateTime[0:space], "-")[0])
	month, _   := strconv.Atoi(strings.Split(dateTime[0:space], "-")[1])
	day, _     := strconv.Atoi(strings.Split(dateTime[0:space], "-")[2])
	hour, _    := strconv.Atoi(strings.Split(dateTime[space + 1:plus - 1], ":")[0])
	minute, _  := strconv.Atoi(strings.Split(dateTime[space + 1:plus - 1], ":")[1])
	second, _  := strconv.Atoi(strings.Split(strings.Split(dateTime[space + 1:plus - 1], ":")[2], ".")[0])
	nanosec, _ := strconv.Atoi(strings.Split(strings.Split(dateTime[space + 1:plus - 1], ":")[2], ".")[1])
	newTime    := time.Date(year, time.Month(month), day, hour, minute, second, nanosec, time.UTC)
	now        := time.Now()
	return newTime.Sub(now).Hours() >= 72
}
