package main

import (
	"context"
	"log"
	"net"

	"github.com/mendoncas/gostify-grpc/pb"
	"google.golang.org/grpc"
)

var (
	songs = []pb.Song{
		{Title: "Frances farmer will have her revenge", Artist: "Nirvana"},
		{Title: "Down in a hole", Artist: "Alice in Chains"},
		{Title: "Doomsday", Artist: "MF DOOM"},
		{Title: "Fullgás", Artist: "Marina Lima"},
		{Title: "Viking hair", Artist: "Dry Cleaning"},
		{Title: "Get Got", Artist: "Death Grips"},
	}
	playlists = []*pb.Playlist{
		{Title: "teste",
			Songs: []*pb.Song{
				&songs[1],
				&songs[2],
				&songs[3],
				&songs[1],
			}},
		{Title: "playlist do ricardo 1", Songs: []*pb.Song{
			&songs[1],
			&songs[2],
			&songs[3],
			&songs[1],
		}},
		{Title: "playlist do ricardo 2"},
		{Title: "playlist do pedro"},
		{Title: "playlist do pedro 2"},
	}
	playlists_do_ricardo = []*pb.Playlist{
		playlists[1],
		playlists[2],
		playlists[0],
	}

	playlists_do_pedro = []*pb.Playlist{
		playlists[3],
		playlists[4],
	}

	ricardo = pb.User{
		Username:  "Ricardo",
		Playlists: playlists_do_ricardo,
	}

	pedro = pb.User{
		Username:  "Pedro",
		Playlists: playlists_do_pedro,
	}
	users = pb.Users{Users: []*pb.User{&pedro, &ricardo}}
)

type Gostify struct {
	pb.UnimplementedGostifyServer
}

func (s *Gostify) GetUsers(c context.Context, in *pb.UserRequest) (*pb.Users, error) {
	return &users, nil
}

func (s *Gostify) GetPlaylists(c context.Context, in *pb.UserRequest) (*pb.Playlists, error) {
	return &pb.Playlists{Playlists: playlists}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGostifyServer(s, &Gostify{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
