package data

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/zabawaba99/firego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/vasialek/VsLinks/helpers"
	"github.com/vasialek/VsLinks/models"
	"google.golang.org/api/option"
)

// LinkRepository provides acces to Link repository as class
type LinkRepository struct {
	db *firego.Firebase
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{}
}

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

// GetLink returns Link by LinkId
func (lr *LinkRepository) GetLink(linkID string) (link models.Link, err error) {
	db, err := lr.getDatabaseClient()
	if err != nil {
		return link, err
	}

	// Get and check if Link exists
	fmt.Printf("  searching Link by LinkId `%s`\n", linkID)
	val := map[string]models.Link{}
	if err = db.Child("links").OrderBy("link_id").EqualTo(linkID).Value(&val); err != nil {
		log.Printf("LinkRepository::GetLink: error getting Link by LinkID. %s\n", err)
		return link, err
	}

	fmt.Println(val)
	for _, v := range val {
		return v, nil
	}

	return link, nil
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

// GetUserLinks returns all links which belongs to user
func (lr *LinkRepository) GetUserLinks(userID string) (links []models.Link, err error) {
	db, err := lr.getDatabaseClient()
	if err != nil {
		return links, err
	}

	fmt.Printf("Searching Link by UserID `%s`\n", userID)
	val := map[string]models.Link{}
	if err = db.Child("links").OrderBy("user_id").EqualTo(userID).Value(&val); err != nil {
		log.Printf("LinkRepository::GetUserLinks: error getting Links by UserID. %s\n", err)
		return links, err
	}

	fmt.Printf("Got %d links for user", len(val))

	for _, v := range val {
		links = append(links, v)
	}

	return links, nil
}

// SetLinkCategory sets LinkCategory for existing link
func (lr *LinkRepository) SetLinkCategory(linkID, categoryID string) error {
	link, err := lr.GetLink(linkID)
	if err != nil {
		log.Printf("LinkRepository::SetLinkCategory: error getting Link. %s\n", err)
		return err
	}

	fmt.Println(link)

	// Ignore DB connection error, we should be connected
	db, _ := lr.getDatabaseClient()
	category := map[string]models.LinkCategory{}
	if err = db.Child("category").OrderBy("link_category_id").EqualTo(categoryID).Value(&category); err != nil {
		log.Printf("LinkRepository::SetLinkCategory: error getting LinkCategory. %s\n", err)
		return err
	}

	fmt.Println(category)

	// Let's change Link category (TODO: validate owner)
	link.LinkCategoryID = categoryID
	log.Printf("Updating Link (ID `%s`) category ID to `%s`\n", link.LinkID, link.LinkCategoryID)
	if err = db.Child("link_test").Child(link.LinkID).Set(&link); err != nil {
		log.Printf("LinkRepository::SetLinkCategory: error updating Link category ID. %s\n", err)
		return err
	}

	return nil
}

func (lr *LinkRepository) validateLinkAndCategory(linkID, categoryID string) (mistakes []string) {
	mistakes = make([]string, 0)

	if len(linkID) != 36 {
		mistakes = append(mistakes, fmt.Sprintf("Length of Link UID should be exactly %d", 36))
	}

	return mistakes
}

func (lr *LinkRepository) getDatabaseClient() (*firego.Firebase, error) {
	if lr.db == nil {
		fmt.Println("Connecting to Firebase...")
		d, err := ioutil.ReadFile("./keys/vsm-links-db-firebase-adminsdk-gubgg-98be377e48.json")
		if err != nil {
			log.Printf("getDatabaseApp: can't read Firebase key: %s\n", err)
			return nil, err
		}

		conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
		if err != nil {
			return nil, err
		}

		lr.db = firego.New("https://vsm-links-db.firebaseio.com/", conf.Client(oauth2.NoContext))
	}
	return lr.db, nil
}

func getDatabaseClient() (db *db.Client, ctx context.Context, err error) {
	c := context.Background()
	opt := option.WithCredentialsFile("./keys/vsm-links-db-firebase-adminsdk-gubgg-98be377e48.json")
	// opt := option.WithCredentialsFile("./keys/vs-links-db-production-726k3-c42e727a62.json")
	config := &firebase.Config{
		DatabaseURL: "https://vsm-links-db.firebaseio.com/",
		// DatabaseURL: "https://vs-links-db.firebaseio.com/",
	}

	app, err := firebase.NewApp(c, config, opt)
	if err != nil {
		return db, ctx, err
	}

	db, err = app.Database(c)
	return db, c, nil
}
