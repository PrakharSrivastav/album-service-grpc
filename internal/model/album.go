package model

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/renstrom/shortuuid"
)

type Album struct {
	Id       string
	Name     string
	ArtistId string
}

func NewAlbum() *Album {
	return &Album{
		Id:       shortuuid.New(),
		Name:     "Some Name",
		ArtistId: shortuuid.New(),
	}
}

func (a *Album) ToProto() *pb.Album {
	return &pb.Album{
		Id:          a.Id,
		Name:        a.Name,
		ArtistId:    a.ArtistId,
		Description: "Album desc",
	}
}
