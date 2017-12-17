var PROTO_PATH = __dirname + '/proto/recording.proto';
var grpc = require('grpc');

if (process.argv.length < 3) {
  console.log('you have to specify a name for your recording!')
  process.exit(1);
}

// Load the protos
var recordingpb = grpc.load(PROTO_PATH).recording;

// Instantiate the client
var client = new recordingpb.RecordingService('server:5050', grpc.credentials.createInsecure());

// Create the request
var recording = new recordingpb.Recording({
  author: new recordingpb.Author({name: 'Nodejs ' + process.version}),
  name: process.argv[2]
});

// Make the call
var call = client.AddRecording(new recordingpb.AddRecordingRequest({recording: recording}), function(error,response){});

// Profit :)
