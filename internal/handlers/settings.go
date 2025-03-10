package handlers

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type SettingsHandlers struct {
	app *pocketbase.PocketBase
}

func NewSettingsHandlers(app *pocketbase.PocketBase) *SettingsHandlers {
	return &SettingsHandlers{
		app: app,
	}
}

func (s *SettingsHandlers) RegisterRoutes(e *core.ServeEvent) error {
	authenticatedOnly := apis.RequireAuth("users")
	e.Router.PUT("/api/settings/credentials", s.HandleChangeCredentials).Bind(authenticatedOnly)
	return nil
}

func (s *SettingsHandlers) HandleChangeCredentials(e *core.RequestEvent) error {
	auth := e.Auth
	if auth == nil {
		return e.UnauthorizedError("User not authenticated.", nil)
	}

	body := &struct {
		Email       string `json:"email" form:"email"`
		Password    string `json:"password" form:"password"`
		OldPassword string `json:"oldPassword" form:"oldPassword"`
	}{}

	if err := e.BindBody(body); err != nil {
		return e.BadRequestError("Failed to read request body.", err)
	}

	collection, err := s.app.FindCollectionByNameOrId("users")
	if err != nil {
		return e.InternalServerError("Failed to find users collection.", err)
	}

	user, err := s.app.FindAuthRecordByEmail(collection, auth.Email())
	if err != nil {
		return e.UnauthorizedError("Not authorized.", nil)
	}

	if !user.ValidatePassword(body.OldPassword) {
		return e.UnauthorizedError("Not authorized.", nil)
	}

	user.SetEmail(body.Email)
	user.SetPassword(body.Password)

	if err := s.app.Save(user); err != nil {
		return e.InternalServerError("Failed to save user.", err)
	}

	return e.NoContent(200)
}
