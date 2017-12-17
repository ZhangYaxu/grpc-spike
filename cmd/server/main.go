package main

import (
	"github.com/rodrigodiez/grpc-spike/recording"
	"google.golang.org/grpc"
	"log"
	"net"
  "os"
  "flag"
)

var delay = flag.Int("delay", 0, "delay in ms between stream elements")

func main() {
	flag.Parse()
	log.Println("listening for gRPC connections on port 5050...")

	// Listen on 5050
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
    handleErr(err)
	}

	// Instantiate a gRPC server
	grpcServer := grpc.NewServer()

	// Register our recording service into the gRPC server
	recording.RegisterRecordingServiceServer(grpcServer, &recordingService{})

	// Profit :)
	grpcServer.Serve(lis)
}

func handleErr(err error){
  log.Fatalf("fatal error", err)
  os.Exit(1)
}
