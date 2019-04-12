package models

// MySettings provides settings for whole project
type MySettings struct {
	IsDevEnvironment bool
	PortToListen     int
	DatabaseURL      string
	DatabaseKey      string
}

// GetEnvironment returns name of environment - DEV or PROD
func (ms *MySettings) GetEnvironment() string {
	if ms.IsDevEnvironment {
		return "DEV"
	}
	return "PROD"
}
