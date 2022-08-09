package db

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"homework/config"
	"homework/internal/storage/models"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	pool, err := pgxpool.Connect(context.Background(), config.DbDSN)
	if err != nil {
		return nil, err
	}

	return &Database{pool: pool}, nil
}

func (s *Database) GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error) {
	query := "SELECT id, title, year FROM public.Movie where id=$1"

	movie := &models.Movie{}
	if err := pgxscan.Get(ctx, s.pool, movie, query, id); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *Database) List(ctx context.Context, limit, offset int) ([]*models.Movie, error) {
	query := "SELECT id, title, year FROM public.Movie ORDER BY id ASC LIMIT $1 OFFSET $2"

	var movies []*models.Movie
	if err := pgxscan.Select(ctx, s.pool, &movies, query, limit, offset); err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *Database) Add(ctx context.Context, m *models.Movie) (uint64, error) {
	var id int
	query := "INSERT INTO public.Movie (title, year) VALUES($1, $2) RETURNING id"
	if err := s.pool.QueryRow(ctx, query, m.Title, m.Year).Scan(&id); err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (s *Database) Update(ctx context.Context, id uint64, newMovie *models.Movie) (error) {
	if _, err := s.GetOneMovie(ctx, id); err != nil {
		return err
	}

	query := "UPDATE public.Movie SET title=$1, year=$2 WHERE id=$3"
	if _, err := s.pool.Exec(ctx, query, newMovie.Title, newMovie.Year, id); err != nil {
		return err
	}
	return nil
}

func (s *Database) Delete(ctx context.Context, id uint64) error {
	query := "DELETE FROM public.Movie WHERE id=$1"
	if _, err := s.pool.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
