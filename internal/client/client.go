package client

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/opentracing/opentracing-go"
)
import "google.golang.org/grpc"

type Client struct {
	artistClientRPC pb.ArtistServiceClient
	trackClientRPC  pb.TrackServiceClient
}

func NewClient(tracer opentracing.Tracer) *Client {
	c := Client{}
	connArtist := getConnection("artist-service-client", "artist-service:6565", tracer)
	connTracks := getConnection("track-service-client", "track-service:6565", tracer)
	c.artistClientRPC = pb.NewArtistServiceClient(connArtist)
	c.trackClientRPC = pb.NewTrackServiceClient(connTracks)
	return &c
}

func getConnection(name string, addr string, tracer opentracing.Tracer) *grpc.ClientConn {
	connection := New(name, addr, tracer)
	return connection.Dial()
}
