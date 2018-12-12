package client

import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
import "google.golang.org/grpc"

type Client struct {
	artistClientRPC pb.ArtistServiceClient
	trackClientRPC  pb.TrackServiceClient
}

func NewClient() *Client {
	c := Client{}
	connArtist := getConnection("artist-service-client", "localhost:6565")//"artist-service:6565")
	connTracks := getConnection("track-service-client", "localhost:6560")//"track-service:6565")
	c.artistClientRPC = pb.NewArtistServiceClient(connArtist)
	c.trackClientRPC = pb.NewTrackServiceClient(connTracks)
	return &c
}

func getConnection(name string, addr string) *grpc.ClientConn {
	connection := New(name, addr)
	return connection.Dial()
}
