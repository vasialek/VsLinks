package helpers

import (
	"fmt"

	"github.com/vasialek/VsLinks/models"
)

// DumpLink outputs all fields to console
func DumpLink(link *models.Link) {
	if link != nil {
		fmt.Println("Link data:")
		fmt.Printf("  LinkID:               %s\n", link.LinkID)
		fmt.Printf("  UserID:               %s\n", link.UserID)
		fmt.Printf("  TypeID:               %d\n", link.TypeID)
		fmt.Printf("  LinkCategoryID:       %s\n", link.LinkCategoryID)
		fmt.Printf("  Title:                %s\n", link.Title)
		fmt.Printf("  Description:          %s\n", link.Description)
		fmt.Printf("  URL:                  %s\n", link.URL)
		fmt.Printf("  Rate:                 %d\n", link.Rate)
		fmt.Printf("  CreatedAt:            %s\n", link.CreatedAt)
	}
}
