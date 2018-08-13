package service

import (
	"context"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/golang/protobuf/ptypes/empty"
)

type Service interface {
	GetAll(*empty.Empty, pb.AlbumService_GetAllServer) error
	GetAlbumByArtist(*pb.SimpleAlbumRequest, pb.AlbumService_GetAlbumByArtistServer) error
	GetAlbumByTrack(context.Context, *pb.SimpleAlbumRequest) (*pb.Album, error)
}
