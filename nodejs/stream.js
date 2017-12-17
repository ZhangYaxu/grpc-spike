var PROTO_PATH = __dirname + '/proto/recording.proto';
var grpc = require('grpc');

// Load the protos
var protoDescriptor = grpc.load(PROTO_PATH);

// Instantiate the client
var client = new protoDescriptor.recording.RecordingService('server:5050', grpc.credentials.createInsecure());

// Make the call
var call = client.ListRecordingsStream(new protoDescriptor.recording.ListRecordingsRequest());
call.on('data', function(recording){
  console.log("ListRecordingsStream() --> '%s' by '%s'", recording.name, recording.author.name);
});

// Profit :)
