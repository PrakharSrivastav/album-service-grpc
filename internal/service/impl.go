package service

import (
	"context"

	"github.com/PrakharSrivastav/album-service-grpc/internal/client"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	repo   *repository
	client *client.Client
}

// Create a new service implementation
func New(db *sqlx.DB) Service {
	return &impl{
		repo:   &repository{db: db},
		client: client.NewClient(),
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
func (f *impl) GetAll() ([]*pb.Album, error) {
	albums, err := f.repo.getAll()
	if err != nil {
		return nil, err
	}

	var protoAlbums []*pb.Album
	for _, album := range albums {
		protoAlbums = append(protoAlbums, album.ToProto())
	}
	return protoAlbums, nil
}

// GetAlbumByArtist gets all the albums for a given artist
func (f *impl) GetAlbumByArtist(artistID string) ([]*pb.Album, error) {
	albums, err := f.repo.getAlbumsByArtist(artistID)
	if err != nil {
		return nil, err
	}

	var protoAlbums []*pb.Album
	for _, album := range albums {
		protoAlbums = append(protoAlbums, album.ToProto())
	}
	return protoAlbums, nil
}

// GetAlbumByTrack get the album for a track
func (f *impl) GetAlbumByTrack(ctx context.Context, req *pb.SimpleAlbumRequest) (*pb.Album, error) {
	trackID := req.GetId()

	track, err := f.client.GetTrack(ctx, trackID)
	if err != nil {
		return nil, err
	}

	album, err := f.repo.get(track.GetAlbumId())
	if err != nil {
		return nil, err
	}
	return album.ToProto(), nil
}

// CleanupAndInit sets up the database with some initial data
func (f *impl) CleanupAndInit() error {
	return f.repo.setupDatabase()
}
