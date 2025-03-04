package migrations

import (
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		email := "admin@admin.com"
		password := "Pass123!"

		if os.Getenv("USER_EMAIL") != "" {
			email = os.Getenv("USER_EMAIL")
		}

		if os.Getenv("USER_PASSWORD") != "" {
			password = os.Getenv("USER_PASSWORD")
		}

		collection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		user := core.NewRecord(collection)
		user.SetEmail(email)
		user.SetPassword(password)
		user.Set("verified", true)
		user.Set("name", "Admin")

		if err := app.Save(user); err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
