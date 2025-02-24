package app

import (
	"licensetogo/internal/handlers"
	licensestorage "licensetogo/internal/license_storage"
	"os"
	"strings"

	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

const DEFAULT_COLLECTION = "license"

type LicenseToGoServer struct {
	app            *pocketbase.PocketBase
	licenseManager *keystogo.Manager
}

func NewLicenseToGoServer() (*LicenseToGoServer, error) {
	app := pocketbase.New()
	// Using default collection name for now, support for multiple license collections
	// will be added in the future
	licenseManager := keystogo.NewManager(licensestorage.NewPocketbaseStorage(DEFAULT_COLLECTION, app))

	return &LicenseToGoServer{
		app:            app,
		licenseManager: licenseManager,
	}, nil
}

func (lts *LicenseToGoServer) Start() error {
	if err := lts.setupMigrations(); err != nil {
		return err
	}

	// Register license handlers
	lts.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// TODO: Handle multiple license collections
		licenseHandler := handlers.NewLicenseHandler(DEFAULT_COLLECTION, lts.licenseManager)
		if err := licenseHandler.RegisterRoutes(e); err != nil {
			return err
		}

		return e.Next()
	})

	return lts.app.Start()
}

func (lts *LicenseToGoServer) setupMigrations() error {
	isGoRun := strings.Contains(strings.Join(os.Args, " "), "--dev")
	return migratecmd.Register(lts.app.App, lts.app.RootCmd, migratecmd.Config{
		Dir:         "./migrations/",
		Automigrate: isGoRun,
	})
}
