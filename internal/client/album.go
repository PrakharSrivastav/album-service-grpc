package client

import (
	"context"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
)

// GetTrack gets the track details by Id
func (c *Client) GetTrack(ctx context.Context, id string) (*pb.Track, error) {
	return c.trackClientRPC.Get(ctx, &pb.SimpleTrackRequest{Id: id})
}
