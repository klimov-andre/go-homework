package facade

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	cachePkg "homework/internal/storage/cache"
	cache "homework/internal/storage/cache/mocks"
	database "homework/internal/storage/db/mocks"
	"homework/internal/storage/models"
	"reflect"
	"testing"
)

func TestStorageFacade_Add(t *testing.T) {
	t.Run("successfully add to db", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		m := &models.Movie {
			Id: 1,
			Title: "Inception",
			Year: 2008,
		}

		db.On("Add", mock.Anything, m).Return(uint64(1), nil)
		cache.On("AddOrUpdate", mock.Anything, uint64(1), m)

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		id, err := storage.Add(context.Background(), m)

		assert.NoError(t, err)
		assert.True(t, id == 1)
		db.AssertNumberOfCalls(t, "Add", 1)
		cache.AssertNumberOfCalls(t, "AddOrUpdate", 1)
	})

	t.Run("error on add to db", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		m := &models.Movie {
			Id: 1,
			Title: "Inception",
			Year: 2008,
		}

		db.On("Add", mock.Anything, m).Return(uint64(0), errors.New("error"))

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		_, err := storage.Add(context.Background(), m)

		assert.Error(t, err)
		db.AssertNumberOfCalls(t, "Add", 1)
		cache.AssertNotCalled(t, "AddOrUpdate")
	})
}

func TestStorageFacade_Update(t *testing.T) {
	t.Run("success db.update", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		m := &models.Movie {
			Id: 1,
			Title: "Inception",
			Year: 2008,
		}

		db.On("Update", mock.Anything,  uint64(1), m).Return(nil)
		cache.On("AddOrUpdate", mock.Anything, uint64(1), m)

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		err := storage.Update(context.Background(), 1, m)

		assert.NoError(t, err)
		db.AssertNumberOfCalls(t, "Update", 1)
		cache.AssertNumberOfCalls(t, "AddOrUpdate", 1)
	})

	t.Run("error on db.update", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		m := &models.Movie {
			Id: 1,
			Title: "Inception",
			Year: 2008,
		}

		db.On("Update", mock.Anything, uint64(1), m).
			Return(errors.New("error"))

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		err := storage.Update(context.Background(), 1, m)

		assert.Error(t, err)
		db.AssertNumberOfCalls(t, "Update", 1)
		cache.AssertNotCalled(t, "AddOrUpdate")
	})
}

func TestStorageFacade_Delete(t *testing.T) {
	t.Run("successfully deleted", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		db.On("Delete", mock.Anything,  uint64(1)).Return(nil)
		cache.On("Delete", mock.Anything, uint64(1)).Return(nil)

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		err := storage.Delete(context.Background(), 1)

		assert.NoError(t, err)
		db.AssertNumberOfCalls(t, "Delete", 1)
		cache.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("error on db.delete", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		db.On("Delete", mock.Anything,  uint64(1)).Return(errors.New("error"))

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		err := storage.Delete(context.Background(), 1)

		assert.Error(t, err)
		db.AssertNumberOfCalls(t, "Delete", 1)
		cache.AssertNotCalled(t, "Delete")
	})

	t.Run("success db.delete, not exists on cache", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		db.On("Delete", mock.Anything,  uint64(1)).Return(nil)
		cache.On("Delete", mock.Anything, uint64(1)).Return(cachePkg.ErrCacheNotExists)

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		err := storage.Delete(context.Background(), 1)

		assert.NoError(t, err)
		db.AssertNumberOfCalls(t, "Delete", 1)
		cache.AssertNumberOfCalls(t, "Delete", 1)
	})
}

func TestStorageFacade_GetOneMovie(t *testing.T) {
	t.Run("successfully find from cache", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		m := &models.Movie {
			Id: 1,
			Title: "Inception",
			Year: 2008,
		}

		cache.On("GetById", mock.Anything, uint64(1)).Return(m, nil)

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		actualMovie, err := storage.GetOneMovie(context.Background(), 1)

		assert.True(t, reflect.DeepEqual(actualMovie, m))
		assert.NoError(t, err)

		cache.AssertNumberOfCalls(t, "GetById", 1)
		db.AssertNotCalled(t, "GetOneMovie")
		db.AssertNotCalled(t, "AddOrUpdate")
	})

	t.Run("no in cache, successfully find from db", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		m := &models.Movie {
			Id: 1,
			Title: "Inception",
			Year: 2008,
		}

		cache.On("GetById", mock.Anything, uint64(1)).
			Return(nil, cachePkg.ErrCacheNotExists)

		db.On("GetOneMovie", mock.Anything, uint64(1)).
			Return(m, nil)

		cache.On("AddOrUpdate", mock.Anything, uint64(1), m)

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		actualMovie, err := storage.GetOneMovie(context.Background(), 1)

		assert.True(t, reflect.DeepEqual(actualMovie, m))
		assert.NoError(t, err)

		db.AssertNumberOfCalls(t, "GetOneMovie", 1)
		cache.AssertNumberOfCalls(t, "GetById", 1)
		cache.AssertNumberOfCalls(t, "AddOrUpdate", 1)
	})

	t.Run("no in cache, error on find from db", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		cache.On("GetById", mock.Anything, uint64(1)).
			Return(nil, cachePkg.ErrCacheNotExists)

		db.On("GetOneMovie", mock.Anything, uint64(1)).
			Return(nil, errors.New("error"))

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		_, err := storage.GetOneMovie(context.Background(), 1)
		assert.Error(t, err)

		db.AssertNumberOfCalls(t, "GetOneMovie", 1)
		cache.AssertNumberOfCalls(t, "GetById", 1)
		cache.AssertNotCalled(t, "AddOrUpdate")
	})
}

func TestStorageFacade_List(t *testing.T) {
	t.Run("success list movies", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		movies := []*models.Movie {
			{
				Id: 1,
				Title: "Inception",
				Year: 2008,
			},
			{
				Id: 2,
				Title: "Hello",
				Year: 2022,
			},
			{
				Id: 3,
				Title: "Goodbye",
				Year: 1999,
			},

		}

		db.On("List", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(movies, nil)
		cache.On("AddOrUpdate", mock.Anything, movies[0].Id, movies[0])
		cache.On("AddOrUpdate", mock.Anything, movies[1].Id, movies[1])
		cache.On("AddOrUpdate", mock.Anything, movies[2].Id, movies[2])

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		actualMovies, err := storage.List(context.Background(), 3, 3, "ASC")

		assert.True(t, reflect.DeepEqual(actualMovies, movies))
		assert.NoError(t, err)

		db.AssertNumberOfCalls(t, "List", 1)
		cache.AssertNumberOfCalls(t, "AddOrUpdate", 3)
	})

	t.Run("database error", func(t *testing.T) {
		db := &database.DatabaseInterface{}
		cache := &cache.CacheInterface{}

		db.On("List", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		storage := storageFacade{
			db:    db,
			cache: cache,
		}

		_, err := storage.List(context.Background(), 3, 3, "ASC")

		assert.Error(t, err)

		db.AssertNumberOfCalls(t, "List", 1)
		cache.AssertNotCalled(t, "AddOrUpdate")
	})
}