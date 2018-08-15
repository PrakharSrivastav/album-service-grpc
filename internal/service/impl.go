package service

import (
	"context"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	repo *repository
}

// Create a new service implementation
func New(db *sqlx.DB) Service {
	return &impl{
		repo: &repository{
			db: db,
		},
	}
}

// Get fetches an album by id
func (f *impl) Get(id string) (*pb.Album, error) {
	album, err := f.repo.get(id)
	if err != nil {
		return nil, err
	}
	return album.ToProto(), nil
}

// GetAll Fetches all the albums in the database
func (f *impl) GetAll(*empty.Empty, pb.AlbumService_GetAllServer) error {
	panic("not implemented")
}

// GetAlbumByArtist gets all the albums for a given artist
func (f *impl) GetAlbumByArtist(*pb.SimpleAlbumRequest, pb.AlbumService_GetAlbumByArtistServer) error {
	panic("not implemented")
}

// GetAlbumByTrack get the album for a track
func (f *impl) GetAlbumByTrack(context.Context, *pb.SimpleAlbumRequest) (*pb.Album, error) {
	panic("not implemented")
}

// CleanupAndInit sets up the database with some initial data
func (f *impl) CleanupAndInit() error {
	return f.repo.setupDatabase()
}
