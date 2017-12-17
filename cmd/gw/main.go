package main

import (
  "net/http"

  "log"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

	"github.com/rodrigodiez/grpc-spike/recording"
)

func main() {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := recording.RegisterRecordingServiceHandlerFromEndpoint(ctx, mux, "server:5050", opts)
  if err != nil {
    log.Fatal(err)
  }

  http.ListenAndServe(":8080", mux)
}
