var PROTO_PATH = __dirname + '/proto/recording.proto';
var grpc = require('grpc');

// Load the protos
var protoDescriptor = grpc.load(PROTO_PATH);

// Instantiate the client
var client = new protoDescriptor.recording.RecordingService('server:5050', grpc.credentials.createInsecure());

// Make the call
client.ListRecordings(new protoDescriptor.recording.ListRecordingsRequest(), function(err, response) {
  for(var i = 0; i < response.recordings.length; i++) {
    console.log("ListRecordings() --> '%s' by '%s'", response.recordings[i].name, response.recordings[i].author.name);
  }
});

// Profit :)
