package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rodrigodiez/grpc-spike/recording"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()
var host = flag.String("host", "127.0.0.1", "gRPC server host")
var port = flag.String("port", "5050", "gRPC server port")
var authorName = flag.String("author", "Lorem ipsum", "recording author name")
var recordingName = flag.String("recording", "Dolor sit amet", "recording name")

func main() {
	flag.Parse()
	log.Printf("%s to gRPC server %s on port %s...", cyan("connecting"), yellow(*host), yellow(*port))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *host, *port), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUserAgent("golang-client"))
	if err != nil {
    handleErr(err)
	}
	defer conn.Close()
	log.Println(green("connected!"))

	client := recording.NewRecordingServiceClient(conn)
	log.Printf("%s '%s' by '%s'", cyan("-->"), yellow(*recordingName), yellow(*authorName))
	client.AddRecording(context.Background(), &recording.Recording{Author: &recording.Author{Name: *authorName}, Name: *recordingName})

	stream, err := client.ListRecordings(context.Background(), &empty.Empty{})
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

		log.Printf("%s '%s' by '%s'", cyan("<--"), yellow(r.Name), yellow(r.Author.Name))
	}
}

func handleErr(err error){
  log.Fatalf("fatal error", err)
  os.Exit(1)
}
