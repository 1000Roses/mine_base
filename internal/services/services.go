package services

import (
	"mine/internal/repositories"
	sample "mine/internal/services/sample"
	"mine/internal/settings"

	"github.com/go-redis/redis"
)

// ! Compose interface
type AppServices struct {
	sample.EventSampleService
}

func NewAppServices(
	appSettings *settings.AppSettings,
	repo *repositories.Repositories,
	rdbCache *redis.Client,
) *AppServices {
	return &AppServices{
		sample.NewSampleService(appSettings, repo),
	}
}
