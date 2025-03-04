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

		if os.Getenv("SUPERUSER_EMAIL") != "" {
			email = os.Getenv("SUPERUSER_EMAIL")
		}

		if os.Getenv("SUPERUSER_PASSWORD") != "" {
			password = os.Getenv("SUPERUSER_PASSWORD")
		}

		superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		record := core.NewRecord(superusers)
		record.Set("email", email)
		record.Set("password", password)

		return app.Save(record)
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
