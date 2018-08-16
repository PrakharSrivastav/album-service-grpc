package model

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
)

// Album represents an entity in database
type Album struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	ArtistID    string `db:"artistId"`
	Description string `db:"description"`
}

// ToProto converts the model to protofub format
func (a *Album) ToProto() *pb.Album {
	return &pb.Album{
		Id:          a.ID,
		Name:        a.Name,
		ArtistId:    a.ArtistID,
		Description: a.Description,
	}
}
