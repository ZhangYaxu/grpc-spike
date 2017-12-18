var PROTO_PATH = __dirname + '/proto/recording.proto';
var grpc = require('grpc');

var recordingpb = grpc.load(PROTO_PATH).recording;

var recordings = [];

function addRecording(call, callback){
  var recording = call.request.recording;
  console.log("AddRecording() <-- '%s' by '%s'", recording.name, recording.author.name,);
  recordings.push(recording);
  callback();
}

function listRecordings(call, callback) {
  console.log("ListRecordings() --> %d recordings listed", recordings.length);
  callback(null, {recordings: recordings});
}

function listRecordingsStream(call) {
  for(var i = 0; i < recordings.length; i++){
    console.log("ListRecordingsStream() --> '%s' by '%s'", recordings[i].name, recordings[i].author.name);
    call.write(recordings[i]);
  }

  call.end();
}


function getServer() {
  var server = new grpc.Server();
  server.addService(recordingpb.RecordingService.service, {
    addRecording: addRecording,
    listRecordings: listRecordings,
    listRecordingsStream: listRecordingsStream
  });

  return server;
}

var server = getServer();
server.bind('0.0.0.0:5050', grpc.ServerCredentials.createInsecure());
console.log("listening for gRPC connections on port 5050...");
server.start();
