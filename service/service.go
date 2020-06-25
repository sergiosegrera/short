package service

import (
	"context"
	"errors"
	"time"

	"github.com/lithammer/shortuuid"
	"github.com/sergiosegrera/short/db"
	"github.com/sergiosegrera/short/models"
	"go.uber.org/zap"
)

type Service interface {
	CreateLink(context.Context, string) (*models.Link, error)
	GetLink(context.Context, string) (*models.Link, error)
}

type ShortService struct {
	DB     db.DB
	Logger *zap.Logger
}

func (s *ShortService) CreateLink(ctx context.Context, url string) (link *models.Link, err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
			"short",
			zap.String("method", "createlink"),
			zap.String("url", url),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	link = &models.Link{}

	// TODO: Verify URL, Max size etc...
	link.Url = url

	link.Id = shortuuid.New()

	err = s.DB.CreateLink(link)
	if err != nil {
		return nil, ErrCreatingLink
	}

	return link, err
}

func (s *ShortService) GetLink(ctx context.Context, id string) (link *models.Link, err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
			"short",
			zap.String("method", "getlink"),
			zap.String("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	link = &models.Link{}
	link, err = s.DB.GetLink(id)
	if err != nil {
		return nil, ErrIdNotFound
	}

	return link, err
}

var (
	ErrInvalidURL   = errors.New("URL is not valid")
	ErrCreatingLink = errors.New("Error creating link")
	ErrIdNotFound   = errors.New("Id not found")
)
