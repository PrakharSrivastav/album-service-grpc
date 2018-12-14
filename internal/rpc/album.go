package rpc

import (
	"context"
	"github.com/PrakharSrivastav/album-service-grpc/internal/service"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// AlbumService implements all the functions defined in protobuf
type AlbumService struct {
	service service.Service
}

// Get fetches an album by Id
func (f *AlbumService) Get(ctx context.Context, req *pb.SimpleAlbumRequest) (*pb.Album, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("service", "album-service-get")
	defer span.Finish()
	return f.service.Get(req.GetId())
}

// GetAll gets list of all the albums
func (f *AlbumService) GetAll(_ *empty.Empty, stream pb.AlbumService_GetAllServer) error {
	ctx := stream.Context()
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("service", "album-service-getAll")
	defer span.Finish()
	albums, err := f.service.GetAll()
	if err != nil {
		return err
	}

	for _, a := range albums {
		if err := stream.Send(a); err != nil {
			return err
		}
	}
	return nil
}

// GetAlbumByArtist gets all the albums for an artist
func (f *AlbumService) GetAlbumByArtist(req *pb.SimpleAlbumRequest, stream pb.AlbumService_GetAlbumByArtistServer) error {
	ctx := stream.Context()
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("service", "album-service-get-album-by-artist")
	defer span.Finish()
	albums, err := f.service.GetAlbumByArtist(req.GetId())
	if err != nil {
		return err
	}

	for _, a := range albums {
		if err := stream.Send(a); err != nil {
			return err
		}
	}
	return nil
}

// GetAlbumByTrack fetches the album for a track
func (f *AlbumService) GetAlbumByTrack(ctx context.Context, req *pb.SimpleAlbumRequest) (*pb.Album, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("service", "album-service-album-by-track")
	defer span.Finish()
	return f.service.GetAlbumByTrack(ctx, req)
}

// Register registers implementation of ArtistService to grpc Server
func (f *AlbumService) Register(server *grpc.Server) {
	pb.RegisterAlbumServiceServer(server, f)
}

// New returns a new instance of AlbumService
func New(service service.Service) *AlbumService {
	return &AlbumService{
		service: service,
	}
}
