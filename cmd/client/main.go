package main

import (
	"flag"
	"fmt"
	recordingpb "github.com/rodrigodiez/grpc-spike/recording"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

var recordingName = flag.String("recording", "Dolor sit amet", "recording name")
var command = flag.String("command", "stream", "service call to run")
var client recordingpb.RecordingServiceClient;

func main() {
	flag.Parse()
	log.Println("connecting to gRPC server 127.0.0.1 on port 5050")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "127.0.0.1", "5050"), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUserAgent("golang-client"))
	if err != nil {
    handleErr(err)
	}
	defer conn.Close()

	client = recordingpb.NewRecordingServiceClient(conn)

	switch *command {
	case "add":
		add();
	case "list":
		list();
	default:
		stream();
	}
}

func add(){
	log.Printf("AddRecording() <-- '%s' by '%s'", *recordingName, "Go client")
	recording := &recordingpb.Recording{Author: &recordingpb.Author{Name: "Go client"}, Name: *recordingName}
	client.AddRecording(context.Background(), &recordingpb.AddRecordingRequest{ Recording: recording})
}

func stream(){
	stream, err := client.ListRecordingsStream(context.Background(), &recordingpb.ListRecordingsRequest{})
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

		log.Printf("ListRecordingsStream() --> '%s' by '%s'", r.Name, r.Author.Name)
	}
}

func list(){
	response, err := client.ListRecordings(context.Background(), &recordingpb.ListRecordingsRequest{})
	if err != nil {
  	handleErr(err)
	}

	for _, r := range response.Recordings {
		log.Printf("ListRecordings() --> '%s' by '%s'", r.Name, r.Author.Name)
	}
}

func handleErr(err error){
  log.Fatalf("fatal error", err)
  os.Exit(1)
}
