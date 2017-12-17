package main

import (
  "github.com/rodrigodiez/grpc-spike/recording"
  "log"
  "time"
  "golang.org/x/net/context"
)

// This is our service
type recordingService struct {
	recordings []*recording.Recording
}

// These service functions implement the interface generated from our .proto file
// - AddRecording
// - ListRecordings
// - ListRecordingsStream
func (s *recordingService) AddRecording(ctx context.Context, r *recording.AddRecordingRequest) (*recording.AddRecordingResponse, error) {
	s.recordings = append(s.recordings, r.Recording)
  log.Printf("AddRecording() <-- '%s' by '%s'", r.Recording.Name, r.Recording.Author.Name)

	return &recording.AddRecordingResponse{}, nil
}

func (s *recordingService) ListRecordings(ctx context.Context, _ *recording.ListRecordingsRequest) (*recording.ListRecordingsResponse, error) {
  log.Printf("ListRecordings() --> %d recordings listed", len(s.recordings))
	return &recording.ListRecordingsResponse{Recordings: s.recordings}, nil
}

func (s *recordingService) ListRecordingsStream(_ *recording.ListRecordingsRequest, stream recording.RecordingService_ListRecordingsStreamServer) error {

	for _, r := range s.recordings {
		time.Sleep(time.Duration(*delay) * time.Millisecond)
    log.Printf("ListRecordingsStream() --> '%s' by '%s'", r.Name, r.Author.Name)

		if err := stream.Send(r); err != nil {
			return err
		}
	}

	return nil
}
