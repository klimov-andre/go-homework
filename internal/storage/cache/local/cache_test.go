package local

import (
	"context"
	"github.com/stretchr/testify/assert"
	cachePkg "homework/internal/storage/cache"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	"reflect"
	"testing"
)

func TestCache_AddOrUpdate(t *testing.T) {
	cache := &cache{
		data:   make(map[uint64]*models.Movie),
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	args := []struct {
		id uint64
		m  models.Movie
	}{
		{
			id: 1,
			m: models.Movie{
				Title: "1",
				Year:  2001,
				Id:  1,
			},
		},
		{
			id: 2,
			m: models.Movie{
				Title: "2",
				Year:  2002,
				Id:  2,
			},
		},
		{
			id: 2,
			m: models.Movie{
				Title: "new 2",
				Year:  2003,
				Id:  2,
			},
		},
	}

	expectedData := map[uint64]*models.Movie {
		1: {
			Title: "1",
			Year:  2001,
			Id:  1,
		},
		2: {
			Title: "new 2",
			Year:  2003,
			Id:  2,
		},
	}

	for i := range args {
		cache.AddOrUpdate(context.Background(), args[i].id, &args[i].m)
	}

	assert.True(t, reflect.DeepEqual(expectedData, cache.data))
}

func TestCache_Delete(t *testing.T) {
	tests := []struct {
		name string
		cache *cache
		argId uint64
		expectedErr error
		expectedData map[uint64]*models.Movie
	} {
		{
			name: "success delete 1 item",
			cache: &cache{
				data:   map[uint64]*models.Movie {
					1: {
						Title: "1",
						Year:  2001,
						Id:  1,
					},
				},
				pool:   connections.NewPool(poolSize, timeoutDuration),
			},
			argId: 1,
			expectedErr: nil,
			expectedData: map[uint64]*models.Movie{},
		},
		{
			name: "delete 1 item not exists",
			cache: &cache{
				data:   map[uint64]*models.Movie {
					1: {
						Title: "1",
						Year:  2001,
						Id:  1,
					},
				},
				pool:   connections.NewPool(poolSize, timeoutDuration),
			},
			argId:       2,
			expectedErr: cachePkg.ErrCacheNotExists,
			expectedData: map[uint64]*models.Movie {
				1: {
					Title: "1",
					Year:  2001,
					Id:  1,
				},
			},
		},
		{
			name: "empty cache delete 1 item not exists",
			cache: &cache{
				data:   map[uint64]*models.Movie {},
				pool:   connections.NewPool(poolSize, timeoutDuration),
			},
			argId:        2,
			expectedErr:  cachePkg.ErrCacheNotExists,
			expectedData: map[uint64]*models.Movie {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.cache.Delete(context.Background(), test.argId)

			if test.expectedErr != nil {
				assert.Error(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.True(t, reflect.DeepEqual(test.expectedData, test.cache.data))
		})
	}
}

func TestCache_GetById(t *testing.T) {
	tests := []struct {
		name string
		cache *cache
		argId uint64
		expectedErr error
		expectedMovie *models.Movie
	} {
		{
			name: "get existed",
			cache: &cache{
				data:   map[uint64]*models.Movie {
					1: {
						Title: "1",
						Year:  2001,
						Id:  1,
					},
				},
				pool:   connections.NewPool(poolSize, timeoutDuration),
			},
			argId: 1,
			expectedErr: nil,
			expectedMovie: &models.Movie{
				Title: "1",
				Year:  2001,
				Id:  1,
			},
		},
		{
			name: "get not existed",
			cache: &cache{
				data:   map[uint64]*models.Movie {
					1: {
						Title: "1",
						Year:  2001,
						Id:  1,
					},
				},
				pool:   connections.NewPool(poolSize, timeoutDuration),
			},
			argId:         2,
			expectedErr:   cachePkg.ErrCacheNotExists,
			expectedMovie: nil,
		},
		{
			name: "empty cache get not existed",
			cache: &cache{
				data:   map[uint64]*models.Movie {},
				pool:   connections.NewPool(poolSize, timeoutDuration),
			},
			argId:         2,
			expectedErr:   cachePkg.ErrCacheNotExists,
			expectedMovie: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualMovie, err := test.cache.GetById(context.Background(), test.argId)

			if test.expectedErr != nil {
				assert.Error(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.True(t, reflect.DeepEqual(test.expectedMovie, actualMovie))
		})
	}
}
