package rpc

import (
	"context"
	"fmt"

	"github.com/PrakharSrivastav/album-service-grpc/internal/model"
	"github.com/PrakharSrivastav/album-service-grpc/internal/service"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type AlbumService struct {
	service service.Service
}

func (f *AlbumService) Get(_ context.Context, req *pb.SimpleAlbumRequest) (*pb.Album, error) {
	return f.service.Get(req.GetId())
}
func (f *AlbumService) GetAll(_ *empty.Empty, stream pb.AlbumService_GetAllServer) error {

	fmt.Println("Inside the function")
	var albums []*model.Album
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	fmt.Println("Added to the list")
	for _, a := range albums {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *AlbumService) GetAlbumByArtist(_ *pb.SimpleAlbumRequest, stream pb.AlbumService_GetAlbumByArtistServer) error {
	fmt.Println("Inside the function")
	var albums []*model.Album
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	albums = append(albums, model.NewAlbum())
	fmt.Println("Added to the list")
	for _, a := range albums {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *AlbumService) GetAlbumByTrack(_ context.Context, stream *pb.SimpleAlbumRequest) (*pb.Album, error) {
	fmt.Println("Inside the function")
	return model.NewAlbum().ToProto(), nil
}

func (f *AlbumService) Register(server *grpc.Server) {
	pb.RegisterAlbumServiceServer(server, f)
}

func New(service service.Service) *AlbumService {
	return &AlbumService{
		service: service,
	}
}
