package handlers

import (
	"strconv"
	"time"

	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/karurosux/keystogo/pkg/models"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type LicenseHandler struct {
	collectionName string
	manager        *keystogo.Manager
}

func NewLicenseHandler(collectionName string, manager *keystogo.Manager) *LicenseHandler {
	return &LicenseHandler{
		collectionName: collectionName,
		manager:        manager,
	}
}

func (lh *LicenseHandler) RegisterRoutes(e *core.ServeEvent) error {
	requiresAuthMiddleware := apis.RequireAuth(lh.collectionName)
	e.Router.GET("/api/"+lh.collectionName, lh.Get).Bind(requiresAuthMiddleware)
	e.Router.POST("/api/"+lh.collectionName, lh.Create)
	return nil
}

func (lh *LicenseHandler) Create(e *core.RequestEvent) error {
	user := e.Auth
	if user == nil {
		return e.UnauthorizedError("User not authenticated.", nil)
	}

	body := &struct {
		Name    string  `json:"name" form:"name"`
		Expires *string `json:"expires" form:"expires"`
	}{}

	if err := e.BindBody(body); err != nil {
		return e.BadRequestError("Failed to read request body.", err)
	}

	var parsedExpires *time.Time
	if body.Expires != nil {
		parsed, err := time.Parse("2006-01-02", *body.Expires)
		if err != nil {
			return e.BadRequestError("Invalid date format. Use YYYY-MM-DD", err)
		}
		parsedExpires = &parsed
	}

	_, key, err := lh.manager.GenerateApiKey(body.Name, nil, nil, parsedExpires)
	if err != nil {
		return e.BadRequestError("Failed to generate API key.", err)
	}

	return e.JSON(200, map[string]string{
		"key": key,
	})
}

func (lh *LicenseHandler) Get(e *core.RequestEvent) error {
	limit := 10
	offset := 0

	if limitStr := e.Request.URL.Query().Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.ParseInt(limitStr, 10, 8)
		if err != nil {
			return e.BadRequestError("Invalid limit parameter", err)
		}
		limit = int(parsedLimit)
	}

	if offsetStr := e.Request.URL.Query().Get("offset"); offsetStr != "" {
		parsedOffset, err := strconv.ParseInt(offsetStr, 10, 8)
		if err != nil {
			return e.BadRequestError("Invalid offset parameter", err)
		}
		offset = int(parsedOffset)
	}

	res, total, err := lh.manager.ListKeys(models.Page{
		Limit:  limit,
		Offset: offset,
	}, models.Filter{})
	if err != nil {
		return e.BadRequestError("", err)
	}

	return e.JSON(200, map[string]any{
		"data":  res,
		"total": total,
	})
}
