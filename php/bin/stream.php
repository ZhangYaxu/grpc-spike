<?php

require __DIR__ . '/../vendor/autoload.php';

// Instantiate the client
$client = new KobaltMusic\Recording\RecordingServiceClient('server:5050', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

// Make the call
$call = $client->ListRecordingsStream(new KobaltMusic\Recording\ListRecordingsRequest());

// Iterate over the stream
$recordings = $call->responses();
foreach ($recordings as $recording) {
  echo sprintf("ListRecordingsStream() --> '%s' by '%s'\n", $recording->getName(), $recording->getAuthor()->getName());
}

// Profit :)
