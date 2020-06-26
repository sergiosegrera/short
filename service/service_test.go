package service

import (
	"context"
	"github.com/sergiosegrera/short/models"
	"go.uber.org/zap"
	"os"
	"testing"
)

var shortService *ShortService
var TEST_URL = "https://google.com"
var TEST_ID string

// Mock DB
type TestDB struct {
	store map[string]string
}

func NewTestDB() *TestDB {
	return &TestDB{
		store: make(map[string]string),
	}
}

func (t *TestDB) CreateLink(link *models.Link) error {
	t.store[link.Id] = link.Url

	return nil
}

func (t *TestDB) GetLink(id string) (*models.Link, error) {
	link := &models.Link{
		Url: t.store[id],
		Id:  id,
	}
	return link, nil
}

// Test endpoints
func TestCreateLink(t *testing.T) {
	link, err := shortService.CreateLink(context.Background(), TEST_URL)

	TEST_ID = link.Id

	if err != nil {
		t.Errorf("CreateLink Failed: %s", err.Error())
	}
}

func TestGetLink(t *testing.T) {
	link, err := shortService.GetLink(context.Background(), TEST_ID)

	if err != nil {
		t.Errorf("GetLink failed: %s", err.Error())
	}

	if link.Url != TEST_URL {
		t.Errorf("GetLink failed: Wanted %s, Got %s", TEST_URL, link.Url)
	}
}

func TestMain(m *testing.M) {
	// Create empty logger
	logger := zap.NewNop()
	// Create service
	shortService = &ShortService{
		Logger: logger,
		DB:     NewTestDB(),
	}

	os.Exit(m.Run())
}
