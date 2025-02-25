package app

import (
	"licensetogo/internal/handlers"
	licensestorage "licensetogo/internal/license_storage"
	"licensetogo/internal/middleware"
	"os"
	"strings"

	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

const (
	DEFAULT_COLLECTION  = "license"
	API_KEYS_COLLECTION = "apikey"
)

type LicenseToGoServer struct {
	app            *pocketbase.PocketBase
	licenseManager *keystogo.Manager
	apiKeyManager  *keystogo.Manager
}

func NewLicenseToGoServer() (*LicenseToGoServer, error) {
	app := pocketbase.New()
	// Using default collection name for now, support for multiple license collections
	// will be added in the future
	licenseManager := keystogo.NewManager(licensestorage.NewPocketbaseStorage(DEFAULT_COLLECTION, app))
	apiKeyManager := keystogo.NewManager(licensestorage.NewPocketbaseStorage(API_KEYS_COLLECTION, app))

	return &LicenseToGoServer{
		app:            app,
		licenseManager: licenseManager,
		apiKeyManager:  apiKeyManager,
	}, nil
}

func (lts *LicenseToGoServer) Start() error {
	if err := lts.setupMigrations(); err != nil {
		return err
	}

	middleware.SetManager(lts.apiKeyManager)

	// Register license handlers
	lts.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// licenses could be requested via endpoint, for that we have the option to generate an api key.
		licenseHandler := handlers.NewLicenseHandler(DEFAULT_COLLECTION, lts.app)
		// Api keys could only be created by users that signin.
		apiKeyHandler := handlers.NewApiKeyHandlers(API_KEYS_COLLECTION, lts.app)

		if err := licenseHandler.RegisterRoutes(e); err != nil {
			return err
		}

		if err := apiKeyHandler.RegisterRoutes(e); err != nil {
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
