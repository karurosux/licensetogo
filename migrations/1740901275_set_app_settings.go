package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		settings.Meta.AppName = "LicenseToGo"
		settings.Logs.MaxDays = 2
		settings.Logs.LogAuthId = true
		settings.Logs.LogIP = false
		settings.RateLimits.Enabled = true
		settings.SMTP.Enabled = false

		return app.Save(settings)
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
