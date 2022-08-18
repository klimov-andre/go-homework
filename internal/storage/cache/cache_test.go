package cache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	"reflect"
	"testing"
)

func TestCache_AddOrUpdate(t *testing.T) {
	cache := NewCache()

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
	cache := &Cache{
		data:   map[uint64]*models.Movie {
			1: {
				Title: "1",
				Year:  2001,
				Id:  1,
			},
			2: {
				Title: "2",
				Year:  2003,
				Id:  2,
			},
		},
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	cache.Delete(context.Background(), )
}
