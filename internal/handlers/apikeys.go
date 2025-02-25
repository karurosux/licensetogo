package handlers

import (
	licensestorage "licensetogo/internal/license_storage"

	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type ApiKeyHandlers struct {
	collectionName string
	manager        *keystogo.Manager
}

func NewApiKeyHandlers(collectionName string, app *pocketbase.PocketBase) *ApiKeyHandlers {
	return &ApiKeyHandlers{
		collectionName: collectionName,
		manager:        keystogo.NewManager(licensestorage.NewPocketbaseStorage(collectionName, app)),
	}
}

func (a *ApiKeyHandlers) RegisterRoutes(e *core.ServeEvent) error {
	requiresAuthMiddleware := apis.RequireAuth("users")
	e.Router.POST("/api/"+a.collectionName, a.Create).Bind(requiresAuthMiddleware)
	return nil
}

func (a *ApiKeyHandlers) Create(e *core.RequestEvent) error {
	user := e.Auth
	if user == nil {
		return e.UnauthorizedError("User not authenticated.", nil)
	}

	body := &struct {
		Name string `json:"name" form:"name"`
	}{}

	if err := e.BindBody(body); err != nil {
		return e.BadRequestError("Failed to read request body.", err)
	}

	_, key, err := a.manager.GenerateApiKey(body.Name, nil, nil, nil)
	if err != nil {
		return e.BadRequestError("Failed to generate API key.", err)
	}

	return e.JSON(200, map[string]string{
		"key": key,
	})
}
