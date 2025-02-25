package licensestorage

import (
	"fmt"
	"strings"

	"github.com/karurosux/keystogo/pkg/keystogo"
	"github.com/karurosux/keystogo/pkg/models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type PocketbaseStorage struct {
	collectionName string
	app            *pocketbase.PocketBase
}

// Clear implements keystogo.Storage.
func (p *PocketbaseStorage) Clear() error {
	fmt.Println("Clearing Pocketbase storage is not enabled.")
	return nil
}

// Create implements keystogo.Storage.
func (p *PocketbaseStorage) Create(apiKey *models.APIKey) error {
	collection, err := p.app.FindCollectionByNameOrId(p.collectionName)
	if err != nil {
		return err
	}

	record := core.NewRecord(collection)
	record.Set("key", apiKey.Key)
	record.Set("name", apiKey.Name)
	if apiKey.ExpiresAt != nil {
		record.Set("expires", apiKey.ExpiresAt)
	}
	record.Set("active", apiKey.Active)
	record.Set("permissions", apiKey.Permissions)
	record.Set("metadata", apiKey.Metadata)

	if err := p.app.Save(record); err != nil {
		return err
	}

	return nil
}

// Delete implements keystogo.Storage.
func (p *PocketbaseStorage) Delete(hashedKey string) error {
	record, err := p.app.FindFirstRecordByData(p.collectionName, "key", hashedKey)
	if err != nil {
		return err
	}

	if err := p.app.Delete(record); err != nil {
		return err
	}

	return nil
}

// Get implements keystogo.Storage.
func (p *PocketbaseStorage) Get(hashedKey string) (*models.APIKey, error) {
	record, err := p.app.FindFirstRecordByData(p.collectionName, "key", hashedKey)
	if err != nil {
		return nil, err
	}

	expires := record.GetDateTime("expires").Time()
	lastused := record.GetDateTime("lastused").Time()
	apiKey := &models.APIKey{
		ID:        record.Id,
		Name:      record.GetString("name"),
		CreatedAt: record.GetDateTime("created").Time(),
		Active:    record.GetBool("active"),
	}

	if !expires.IsZero() {
		apiKey.ExpiresAt = &expires
	}

	if !lastused.IsZero() {
		apiKey.LastUsedAt = &lastused
	}

	if err := record.UnmarshalJSONField("permissions", &apiKey.Permissions); err != nil {
		return nil, err
	}

	if err := record.UnmarshalJSONField("metadata", &apiKey.Metadata); err != nil {
		return nil, err
	}

	return apiKey, nil
}

// List implements keystogo.Storage.
func (p *PocketbaseStorage) List(page models.Page, filter models.Filter) ([]models.APIKey, int64, error) {
	result := []models.APIKey{}
	f := []string{}
	params := dbx.Params{}

	if filter.Active != nil {
		params["active"] = *filter.Active
		f = append(f, "active = {:active}")
	}
	if filter.Name != nil {
		params["name"] = *filter.Name
		f = append(f, "name = {:name}")
	}

	records, err := p.app.FindRecordsByFilter(p.collectionName, strings.Join(f, " && "), "-created", page.Limit, page.Offset, params)
	if err != nil {
		return nil, 0, err
	}

	total, err := p.app.CountRecords(p.collectionName, dbx.NewExp(strings.Join(f, " && "), params))
	if err != nil {
		return nil, 0, err
	}

	for index := range records {
		key := models.APIKey{}
		curr := records[index]
		curr.Collection().RawOptions.Scan(&key)
		result = append(result, key)
	}

	return result, total, nil
}

// Ping implements keystogo.Storage.
func (p *PocketbaseStorage) Ping() error {
	return nil
}

// Update implements keystogo.Storage.
func (p *PocketbaseStorage) Update(hashKey string, update models.ApiKeyUpdate) error {
	record, err := p.app.FindFirstRecordByData(p.collectionName, "key", hashKey)
	if err != nil {
		return err
	}

	if update.Name != nil {
		record.Set("name", *update.Name)
	}
	if update.ExpiresAt != nil {
		record.Set("expires", *update.ExpiresAt)
	}
	if update.LastUsedAt != nil {
		record.Set("lastused", *update.LastUsedAt)
	}
	if update.Active != nil {
		record.Set("active", update.Active)
	}
	if update.Permissions != nil {
		record.Set("permissions", update.Permissions)
	}
	if update.Metadata != nil {
		record.Set("metadata", update.Metadata)
	}

	if err := p.app.Save(record); err != nil {
		return err
	}

	return nil
}

func NewPocketbaseStorage(collectioName string, app *pocketbase.PocketBase) keystogo.Storage {
	return &PocketbaseStorage{
		collectionName: collectioName,
		app:            app,
	}
}
