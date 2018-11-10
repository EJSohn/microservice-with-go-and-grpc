package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/golang/protobuf/proto"

	"grpc-tutorial/grpc-polygot/testdata"
	pb "grpc-tutorial/grpc-polygot/api"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/user_db.json", "A json file containing a list of users")
	port       = flag.Int("port", 10000, "The server port")
)

type Server struct {
	savedUsers []*pb.User // read-only after initialized
	mu			sync.Mutex // protects routeNotes
}

func (s *Server) GetPoint(ctx context.Context, profile *pb.Profile) (*pb.Point, error) {
	for _, user := range s.savedUsers {
		if profile.Age == user.Profile.Age && profile.Name == user.Profile.Name {
			return user.Cache, nil
		}
	}
	return nil, errors.New("error: there is no such user.")
}

func (s *Server) ListUsers(point *pb.Point, stream pb.UserGuide_ListUsersServer) error {
	for _, user := range s.savedUsers {
		if proto.Equal(user.Cache, point) {
			if err := stream.Send(user); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Server) loadUsers(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	if err := json.Unmarshal(file, &s.savedUsers); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func newServer() *Server {
	s := &Server{}
	s.loadUsers(*jsonDBFile)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		log.Printf("Establishing TLS...")
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserGuideServer(grpcServer, newServer())
	
	log.Printf("gRPC Server Running...")
	grpcServer.Serve(lis)
}