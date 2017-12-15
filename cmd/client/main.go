package main

import (
	"flag"
	"fmt"
	"github.com/rodrigodiez/grpc-spike/recording"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

var host = flag.String("host", "127.0.0.1", "gRPC server host")
var port = flag.String("port", "5050", "gRPC server port")
var authorName = flag.String("author", "Lorem ipsum", "recording author name")
var recordingName = flag.String("recording", "Dolor sit amet", "recording name")
var command = flag.String("command", "stream", "service call to run")
var client recording.RecordingServiceClient;

func main() {
	flag.Parse()
	log.Printf("connecting to gRPC server %s on port %s...", *host, *port)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *host, *port), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUserAgent("golang-client"))
	if err != nil {
    handleErr(err)
	}
	defer conn.Close()

	client = recording.NewRecordingServiceClient(conn)

	switch *command {
	case "add":
		add();
	default:
		stream();
	}
}

func stream(){
	stream, err := client.ListRecordingsStream(context.Background(), &recording.KBEmpty{})
	if err != nil {
    handleErr(err)
	}

	for {
		r, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
      handleErr(err)
		}

		log.Printf("received recording '%s' by '%s' from %s", r.Name, r.Author.Name, *host)
	}
}

func add(){
	log.Printf("sending recording '%s' by '%s' to %s", *recordingName, *authorName, *host)
	client.AddRecording(context.Background(), &recording.Recording{Author: &recording.Author{Name: *authorName}, Name: *recordingName})
}

func handleErr(err error){
  log.Fatalf("fatal error", err)
  os.Exit(1)
}
