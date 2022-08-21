package storage

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	apiPkg "homework/internal/api/storage/server"
	"homework/internal/storage/facade"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
	"net"
	"testing"
)

const (
	DbConnectionString = "host=127.0.0.1 port=5556 user=user password=password dbname=movies-test sslmode=disable"
	bufSize = 1024 * 1024
)


func runTestStorage(t *testing.T) *bufconn.Listener {
	storage, err := facade.NewStorage(DbConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	listener := bufconn.Listen(bufSize)
	server := grpc.NewServer()
	pb.RegisterStorageServer(server, apiPkg.New(storage))
	go func() {
		if err = server.Serve(listener); err != nil {
			t.Fatal(err)
		}
	}()
	return listener
}

type clientStorageFixture struct {
	conn *grpc.ClientConn
	client pb.StorageClient
}

func setupClientStorage(t *testing.T, listener *bufconn.Listener) *clientStorageFixture {
	ctx := context.Background()

	dialer := func (context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatal(err)
	}

	return &clientStorageFixture{
		conn: conn,
		client: pb.NewStorageClient(conn),
	}
}

func (c *clientStorageFixture) TearDown() {
	c.conn.Close()
}

func oneMoviePrepare(movie *models.Movie, storage *clientStorageFixture) (uint64, error) {
	response, err := storage.client.MovieCreate(context.Background(), &pb.StorageMovieCreateRequest{
		Title: movie.Title,
		Year: int32(movie.Year),
	})

	return response.GetId(), err
}

func movieListPrepare(storage *clientStorageFixture, movies ...models.Movie) ([]uint64, error) {
	var res []uint64

	for _, m := range movies {
		response, err := storage.client.MovieCreate(context.Background(), &pb.StorageMovieCreateRequest{
			Title: m.Title,
			Year: int32(m.Year),
		})

		if err != nil {
			return nil, err
		}

		res = append(res, response.GetId())
	}


	return res, nil
}