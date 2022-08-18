package db

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"homework/config"
	"homework/internal/storage/models"
)

//go:generate mockery --name=DatabaseInterface --case=snake --with-expecter --structname=DatabaseInterface --exported=true

// DatabaseInterface is a common interface for communication with database
type DatabaseInterface interface {
	GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error)
	List(ctx context.Context, limit, offset int, order string) ([]*models.Movie, error)
	Add(ctx context.Context, m *models.Movie) (uint64, error)
	Update(ctx context.Context, id uint64, newMovie *models.Movie) error
	Delete(ctx context.Context, id uint64) error
}

// PgxInterface encapsulates pgxpool.Pool behavior
// This interface is need to simulate pgx behavior for testing
type PgxInterface interface {
	Begin(context.Context) (pgx.Tx, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// Database implements DatabaseInterface
type Database struct {
	pool PgxInterface
}

func NewDatabase() (DatabaseInterface, error) {
	pool, err := pgxpool.Connect(context.Background(), config.DbDSN)
	if err != nil {
		return nil, err
	}

	return &Database{pool: pool}, nil
}

var (
	selectOneMovieQuery = "SELECT id, title, year FROM public.Movie where id=$1"
	selectMovieListQueryFmt = "SELECT id, title, year FROM public.Movie ORDER BY id %s LIMIT $1 OFFSET $2"
	insertOneMovieQuery = "INSERT INTO public.Movie (title, year) VALUES($1, $2) RETURNING id"
	updateOneMovieQuery = "UPDATE public.Movie SET title=$1, year=$2 WHERE id=$3"
	deleteOneMovieQuery = "DELETE FROM public.Movie WHERE id=$1"
)

// GetOneMovie returns one movie with specified id from database
func (s *Database) GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error) {
	query := selectOneMovieQuery

	movie := &models.Movie{}
	if err := pgxscan.Get(ctx, s.pool, movie, query, int(id)); err != nil {
		return nil, err
	}

	return movie, nil
}

// List returns list of movies from database
func (s *Database) List(ctx context.Context, limit, offset int, order string) ([]*models.Movie, error) {
	query := fmt.Sprintf(selectMovieListQueryFmt, order)

	var movies []*models.Movie
	if err := pgxscan.Select(ctx, s.pool, &movies, query, limit, offset); err != nil {
		return nil, err
	}

	return movies, nil
}

// Add adds one movie to database, returns id of new movie on success
func (s *Database) Add(ctx context.Context, m *models.Movie) (uint64, error) {
	var id int
	query := insertOneMovieQuery
	if err := s.pool.QueryRow(ctx, query, m.Title, m.Year).Scan(&id); err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// Update change movie data with specified id
func (s *Database) Update(ctx context.Context, id uint64, newMovie *models.Movie) error {
	if _, err := s.GetOneMovie(ctx, id); err != nil {
		return err
	}

	query := updateOneMovieQuery
	if _, err := s.pool.Exec(ctx, query, newMovie.Title, newMovie.Year, id); err != nil {
		return err
	}
	return nil
}

// Delete removes one movie from database
func (s *Database) Delete(ctx context.Context, id uint64) error {
	query := deleteOneMovieQuery
	if _, err := s.pool.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
