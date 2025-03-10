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

	lts.overrideUI()

	// Register license handlers
	lts.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// licenses could be requested via endpoint, for that we have the option to generate an api key.
		licenseHandler := handlers.NewLicenseHandler(DEFAULT_COLLECTION, lts.app)
		// Api keys could only be created by users that signin.
		apiKeyHandler := handlers.NewApiKeyHandlers(API_KEYS_COLLECTION, lts.app)
		settingsHandler := handlers.NewSettingsHandlers(lts.app)

		if err := licenseHandler.RegisterRoutes(e); err != nil {
			return err
		}

		if err := apiKeyHandler.RegisterRoutes(e); err != nil {
			return err
		}

		if err := settingsHandler.RegisterRoutes(e); err != nil {
			return err
		}

		return e.Next()
	})

	return lts.app.Start()
}

func (lts *LicenseToGoServer) isDev() bool {
	return strings.Contains(strings.Join(os.Args, " "), "--dev")
}

func (lts *LicenseToGoServer) setupMigrations() error {
	isGoRun := lts.isDev()
	return migratecmd.Register(lts.app.App, lts.app.RootCmd, migratecmd.Config{
		Dir:         "./migrations/",
		Automigrate: isGoRun,
	})
}

func (lts *LicenseToGoServer) overrideUI() {
	if lts.isDev() {
		return
	}

	lts.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		e.Router.BindFunc(func(r *core.RequestEvent) error {
			if strings.HasPrefix(r.Request.URL.Path, "/_/") {
				return r.BadRequestError("", nil)
			}
			return r.Next()
		})
		return e.Next()
	})
}
