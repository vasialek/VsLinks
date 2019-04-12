// +build !heroku

package models

// Settings specific project settings in DEV or PROD
var Settings = MySettings{
	IsDevEnvironment: true,
	PortToListen:     9791,
	DatabaseURL:      "https://vsm-links-db.firebaseio.com",
	DatabaseKey:      "./keys/vsm-links-db-firebase-adminsdk-gubgg-98be377e48.json",
}
