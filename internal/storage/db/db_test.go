package db

import (
	"context"
	"fmt"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"homework/internal/storage/models"
	"regexp"
	"testing"
)

func setupDatabase(t *testing.T) (*Database, pgxmock.PgxConnIface) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return &Database{
		pool: mock,
	}, mock
}

func TestGetOneMovie(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		db, mock := setupDatabase(t)
		defer mock.Close(context.Background())

		movieID := uint64(1)
		rows := pgxmock.NewRows([]string{"id", "title", "year"}).AddRow(movieID, "Hello", 2002)
		ctx := context.Background()
		mock.ExpectQuery(regexp.QuoteMeta(selectOneMovieQuery)).
			WithArgs(int(movieID)).
			WillReturnRows(
			rows,
		)

		expectedMovie := &models.Movie{
			Id:    movieID,
			Title: "Hello",
			Year:  2002,
		}

		actualMovie, err := db.GetOneMovie(ctx, movieID)


		assert.NoError(t, err)
		assert.Equal(t, expectedMovie, actualMovie)
		if mockErr := mock.ExpectationsWereMet(); mockErr != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("unsuccessful case", func(t *testing.T) {
		db, mock := setupDatabase(t)
		defer mock.Close(context.Background())

		mock.ExpectQuery(regexp.QuoteMeta(insertOneMovieQuery)).
			WithArgs("", 0).WillReturnError(fmt.Errorf("some error"))

		ctx := context.Background()
		_, err := db.Add(ctx, &models.Movie{
			Id:    0,
			Title: "",
			Year:  0,
		})

		assert.Error(t, err)
		if mockErr := mock.ExpectationsWereMet(); mockErr != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("success case", func(t *testing.T) {
		db, mock := setupDatabase(t)
		defer mock.Close(context.Background())

		expectedId := uint64(1)
		mock.ExpectQuery(regexp.QuoteMeta(insertOneMovieQuery)).
			WithArgs("Hello", 2002).WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))

		ctx := context.Background()
		actualID, err := db.Add(ctx, &models.Movie{
			Title: "Hello",
			Year:  2002,
		})

		assert.NoError(t, err)
		assert.Equal(t, expectedId, actualID)
		if mockErr := mock.ExpectationsWereMet(); mockErr != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
