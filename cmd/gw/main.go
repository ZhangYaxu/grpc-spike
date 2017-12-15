package main

import (
  "flag"
  "net/http"

  "github.com/golang/glog"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

	"github.com/rodrigodiez/grpc-spike/recording"
  "fmt"
)

var (
  host = flag.String("host", "127.0.0.1", "gRPC server host")
  port = flag.String("port", "5050", "gRPC server port")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := recording.RegisterRecordingServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", *host, *port), opts)
  if err != nil {
    return err
  }

  return http.ListenAndServe(":8080", mux)
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}
