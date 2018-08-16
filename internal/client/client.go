package client

import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
import "google.golang.org/grpc"

type Client struct {
	artistClientRpc pb.ArtistServiceClient
	trackClientRpc  pb.TrackServiceClient
}

func NewClient() *Client {
	c := Client{}
	connArtist := getConnection("artist-service-client", "artist-service:6565")
	connTracks := getConnection("track-service-client", "track-service:6565")
	c.artistClientRpc = pb.NewArtistServiceClient(connArtist)
	c.trackClientRpc = pb.NewTrackServiceClient(connTracks)
	return &c
}

func getConnection(name string, addr string) *grpc.ClientConn {
	connection := New(name, addr)
	return connection.Dial()
}
