// +build heroku

package models

var Settings = MySettings{
	IsDevEnvironment: false,
	DatabaseURL:      "https://vsm-links-db.firebaseio.com",
	DatabaseKey:      "./keys/vsm-links-db-firebase-adminsdk-gubgg-98be377e48.json",
}
