package main

import (
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rodrigodiez/grpc-spike/recording"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
  "os"
  "fmt"
  "flag"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()
var port = flag.String("port", "5050", "gRPC server port")

type recordingService struct {
	recordings []*recording.Recording
}

func (s *recordingService) AddRecording(ctx context.Context, r *recording.Recording) (*empty.Empty, error) {
	s.recordings = append(s.recordings, r)
  md, _ := metadata.FromIncomingContext(ctx);
  log.Printf("%s '%s' by '%s'", cyan("<--"), yellow(r.Name), yellow(r.Author.Name))
  log.Printf("%s", md["user-agent"]);

	return &empty.Empty{}, nil
}

func (s *recordingService) ListRecordings(_ *empty.Empty, stream recording.RecordingService_ListRecordingsServer) error {
	log.Printf("listing %d recordings", len(s.recordings))
	for _, r := range s.recordings {
		if err := stream.Send(r); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.Printf("%s for gRPC connections on port %s...", cyan("listening"), yellow(*port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
    handleErr(err)
	}
	grpcServer := grpc.NewServer()
	recording.RegisterRecordingServiceServer(grpcServer, &recordingService{})
	grpcServer.Serve(lis)
}

func handleErr(err error){
  log.Fatalf("fatal error", err)
  os.Exit(1)
}
