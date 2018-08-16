package service

import (
	"context"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
)

type Service interface {
	Get(id string) (*pb.Album, error)
	GetAll() ([]*pb.Album, error)
	GetAlbumByArtist(album string) ([]*pb.Album, error)
	GetAlbumByTrack(context.Context, *pb.SimpleAlbumRequest) (*pb.Album, error)
	CleanupAndInit() error
}
