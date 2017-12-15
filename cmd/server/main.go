package main

import (
	"github.com/rodrigodiez/grpc-spike/recording"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
  "os"
  "fmt"
  "flag"
	"time"
)

var port = flag.String("port", "5050", "gRPC server port")
var delay = flag.Int("delay", 0, "delay in ms between stream elements")

type recordingService struct {
	recordings []*recording.Recording
}

func (s *recordingService) AddRecording(ctx context.Context, r *recording.Recording) (*recording.KBEmpty, error) {
	s.recordings = append(s.recordings, r)
  md, _ := metadata.FromIncomingContext(ctx);
  log.Printf("received recording '%s' by '%s' from %s", r.Name, r.Author.Name, md["user-agent"])

	return &recording.KBEmpty{}, nil
}

func (s *recordingService) ListRecordings(ctx context.Context, _ *recording.KBEmpty) (*recording.ListRecordingsResponse, error) {
  md, _ := metadata.FromIncomingContext(ctx);
	log.Printf("%s recordings sent to %s", len(s.recordings), md["user-agent"])

	return &recording.ListRecordingsResponse{Recordings: s.recordings}, nil
}

func (s *recordingService) ListRecordingsStream(_ *recording.KBEmpty, stream recording.RecordingService_ListRecordingsStreamServer) error {
  md, _ := metadata.FromIncomingContext(stream.Context());
	log.Printf("%d recordings streamed to %s", len(s.recordings), "streamed", md["user-agent"])

	for _, r := range s.recordings {
		time.Sleep(time.Duration(*delay) * time.Millisecond)
		if err := stream.Send(r); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()
	log.Printf("listening for gRPC connections on port %s...", *port)
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
