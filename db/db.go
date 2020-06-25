package db

import "github.com/sergiosegrera/short/models"

type DB interface {
	CreateLink(*models.Link) error
	GetLink(string) (*models.Link, error)
}
