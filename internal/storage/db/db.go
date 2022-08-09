package db

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"homework/config"
	storagePkg "homework/internal/storage"
	"homework/internal/storage/models"
)

var _ storagePkg.Storage = (*Database)(nil)

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

func (s *Database) GetOneMovie(id uint64) (*models.Movie, error) {
	query := "SELECT id, title, year FROM public.Movie where id=$1"

	var movie []*models.Movie
	if err := pgxscan.Select(context.Background(), s.pool, &movie, query, id); err != nil {
		return nil, err
	}

	if len(movie) == 0 {
		return nil, storagePkg.ErrMovieNotExists
	}

	return movie[0], nil
}

func (s *Database) List(limit, offset int) ([]*models.Movie, error) {
	query := "SELECT id, title, year FROM public.Movie ORDER BY id ASC LIMIT $1 OFFSET $2"

	var movies []*models.Movie
	if err := pgxscan.Select(context.Background(), s.pool, &movies, query, limit, offset); err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *Database) Add(m *models.Movie) error {
	var id int
	query := "INSERT INTO public.Movie (title, year) VALUES($1, $2) RETURNING id"
	if err := s.pool.QueryRow(context.Background(), query, m.Title, m.Year).Scan(&id); err != nil {
		return err
	}
	m.SetId(uint64(id))
	return nil
}

func (s *Database) Update(id uint64, newMovie *models.Movie) (*models.Movie, error) {
	if _, err := s.GetOneMovie(id); err != nil {
		return nil, err
	}

	query := "UPDATE public.Movie SET title=$1, year=$2 WHERE id=$3"
	if _, err := s.pool.Exec(context.Background(), query, newMovie.Title, newMovie.Year, id); err != nil {
		return nil, err
	}
	newMovie.SetId(id)
	return newMovie, nil
}

func (s *Database) Delete(id uint64) error {
	if _, err := s.GetOneMovie(id); err != nil {
		return err
	}

	query := "DELETE FROM public.Movie WHERE id=$1"
	if _, err := s.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}

	return nil
}
