package handlers

import (
	licensestorage "licensetogo/internal/license_storage"
	"licensetogo/internal/middleware"
	"strconv"
	"time"

	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/karurosux/keystogo/pkg/models"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type LicenseHandler struct {
	collectionName       string
	manager              *keystogo.Manager
	protectionMiddleware func(e *core.RequestEvent) error
}

func NewLicenseHandler(collectionName string, app *pocketbase.PocketBase) *LicenseHandler {
	m := keystogo.NewManager(licensestorage.NewPocketbaseStorage(collectionName, app))
	pm := middleware.ApiKeyOrUserMiddleware()
	return &LicenseHandler{
		collectionName:       collectionName,
		manager:              m,
		protectionMiddleware: pm,
	}
}

func (lh *LicenseHandler) RegisterRoutes(e *core.ServeEvent) error {
	e.Router.GET("/api/"+lh.collectionName, lh.Get).BindFunc(lh.protectionMiddleware)
	e.Router.POST("/api/"+lh.collectionName, lh.Create).BindFunc(lh.protectionMiddleware)
	e.Router.POST("/api/"+lh.collectionName+"/validate", lh.Validate)
	return nil
}

func (lh *LicenseHandler) Validate(e *core.RequestEvent) error {
	body := &struct {
		Key         string              `json:"key" form:"key"`
		Permissions []models.Permission `json:"permissions" form:"permissions"`
	}{}

	if err := e.BindBody(body); err != nil {
		return e.BadRequestError("Failed to read request body.", err)
	}

	res := lh.manager.ValidateKey(body.Key, body.Permissions)
	return e.JSON(200, res)
}

func (lh *LicenseHandler) Create(e *core.RequestEvent) error {
	body := &struct {
		Name        string               `json:"name" form:"name"`
		Expires     *string              `json:"expires" form:"expires"`
		Permissions *[]models.Permission `json:"permissions" form:"permissions"`
		Metadata    *map[string]any      `json:"metadata" form:"metadata"`
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

	_, key, err := lh.manager.GenerateApiKey(body.Name, body.Permissions, body.Metadata, parsedExpires)
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
