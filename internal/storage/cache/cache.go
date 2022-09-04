package cache

import (
	"context"
	"homework/internal/storage/models"
)

//go:generate mockery --name=CacheInterface --case=snake --with-expecter --structname=CacheInterface --exported=true

type CacheInterface interface{
	AddOrUpdate(ctx context.Context, id uint64, m *models.Movie)
	Delete(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*models.Movie, error)
}


type Type int
const (
	TypeLocal Type = iota
	TypeRedis
)
