package omfhelpers

import(
    "math/rand"
    "time"
)
// from http://golangcookbook.blogspot.com/2012/11/generate-random-number-in-given-range.html
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func GetAltitude() int{
	return random(1, 40) * 1000
}

func GetHeading() int{
	return random(1, 36) * 10;
}

func GetCurrentTimeFormatted() string {
	// return current time in the OMF-specified format

	//from https://golang.org/pkg/time/ and https://stackoverflow.com/questions/44873825/how-to-get-timestamp-of-utc-time-with-golang
	curr := time.Now().UTC()
	formatted := curr.Format(time.RFC3339)
	return formatted
}