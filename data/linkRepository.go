package data

import (
	"context"
	"fmt"

	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/vasialek/VsLinks/helpers"
	"github.com/vasialek/VsLinks/models"
	"google.golang.org/api/option"
)

// CreateLink creates new Link in database
func CreateLink(link models.Link) error {
	c, ctx, err := getDatabaseClient()
	if err != nil {
		return err
	}

	uid := helpers.GetUID()
	fmt.Printf("  saving Link under /links/%s\n", uid)
	link.LinkID = uid
	err = c.NewRef("/links/"+uid).Set(ctx, link)
	if err != nil {
		return err
	}

	return nil
}

// GetAllLinks returns all links in database
func GetAllLinks() ([]models.Link, error) {
	c, ctx, err := getDatabaseClient()
	if err != nil {
		return nil, err
	}

	var data map[string]models.Link

	err = c.NewRef("/links").Get(ctx, &data)
	if err != nil {
		return nil, err
	}

	pos := 0
	list := make([]models.Link, len(data))
	for _, v := range data {
		list[pos] = v
		pos++
	}

	return list, nil
}

func getDatabaseClient() (db *db.Client, ctx context.Context, err error) {
	c := context.Background()
	opt := option.WithCredentialsFile("./keys/vsm-links-db-firebase-adminsdk-gubgg-98be377e48.json")
	config := &firebase.Config{
		DatabaseURL: "https://vsm-links-db.firebaseio.com/",
	}

	app, err := firebase.NewApp(c, config, opt)
	if err != nil {
		return db, ctx, err
	}

	db, err = app.Database(c)
	return db, c, nil
}
