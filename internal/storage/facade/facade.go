package facade

import (
	"context"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	storageCfg "homework/config/storage"
	"homework/internal/api/storage/metrics"
	cachePkg "homework/internal/storage/cache"
	"homework/internal/storage/cache/local"
	"homework/internal/storage/cache/redis"
	dbPkg "homework/internal/storage/db"
	"homework/internal/storage/models"
)

var _ StorageFacade = (*storageFacade)(nil)

//go:generate mockery --name=StorageFacade --case=snake --with-expecter --structname=StorageFacade --exported=true

type StorageFacade interface {
	List(ctx context.Context, limit, offset int, order string) ([]*models.Movie, error)
	Add(ctx context.Context, m *models.Movie) (uint64, error)
	Update(ctx context.Context, id uint64, newMovie *models.Movie) error
	Delete(ctx context.Context, id uint64) error
	GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error)
}

type storageFacade struct {
	db dbPkg.DatabaseInterface
	cache cachePkg.CacheInterface
}

func NewStorage(dbConnection string, cacheType cachePkg.Type) (StorageFacade, error) {
	db, err := dbPkg.NewDatabase(dbConnection)
	if err != nil {
		return nil, errors.Wrap(err, "could not init database")
	}

	var cache cachePkg.CacheInterface

	switch cacheType {
	case cachePkg.TypeLocal:
		cache = local.NewLocalCache()
	case cachePkg.TypeRedis:
		cache = redis.NewRedisCache(metrics.CacheMissCounter, metrics.CacheHitCounter)
	}

	return &storageFacade{
		db: db,
		cache: cache,
	}, err
}

func (s *storageFacade) List(ctx context.Context, limit, offset int, order string) ([]*models.Movie, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "List")
	defer span.End()

	movies, err := s.db.List(ctx, limit, offset, order)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	for _, m := range movies {
		s.cache.AddOrUpdate(ctx, m.Id, m)
	}

	return movies, nil
}

func (s *storageFacade) Add(ctx context.Context, m *models.Movie) (uint64, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "Add")
	defer span.End()

	id, err := s.db.Add(ctx, m)
	if err != nil {
		span.RecordError(err)
		return 0, err
	}

	s.cache.AddOrUpdate(ctx, id, m)
	return id, err
}

func (s *storageFacade) Update(ctx context.Context, id uint64, newMovie *models.Movie) error {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "Update")
	defer span.End()

	if err := s.db.Update(ctx, id, newMovie); err != nil {
		span.RecordError(err)
		return err
	}

	s.cache.AddOrUpdate(ctx, id, newMovie)
	return nil
}

func (s *storageFacade) Delete(ctx context.Context, id uint64) error {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "Delete")
	defer span.End()

	if err := s.db.Delete(ctx, id); err != nil {
		span.RecordError(err)
		return err
	}

	if err := s.cache.Delete(ctx, id); err != nil {
		if !errors.Is(err, cachePkg.ErrCacheNotExists) {
			span.RecordError(err)
			return err
		}
	}
	return nil
}

func (s *storageFacade) GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "GetOneMovie")
	defer span.End()

	m, err := s.cache.GetById(ctx, id)
	if err != nil && !errors.Is(err, cachePkg.ErrCacheNotExists) {
		span.RecordError(err)
		return nil, err
	}
	if m != nil {
		return m, nil
	}

	m, err = s.db.GetOneMovie(ctx, id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	s.cache.AddOrUpdate(ctx, id, m)

	return m, nil
}