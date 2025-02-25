package middleware

import (
	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/pocketbase/pocketbase/core"
)

var manager *keystogo.Manager

func SetManager(m *keystogo.Manager) {
	manager = m
}

func ApiKeyOrUserMiddleware() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		// Check if user is authenticated
		auth := e.Auth
		if auth != nil {
			return e.Next()
		}

		// Check  if license key is present
		key := e.Request.Header.Get("apikey")
		if key == "" {
			return e.UnauthorizedError("", nil)
		}

		res := manager.ValidateKey(key, nil)
		if res.Error != nil || !res.Valid {
			return e.UnauthorizedError("Ivalid api key.", res.Error)
		}

		return e.Next()
	}
}
