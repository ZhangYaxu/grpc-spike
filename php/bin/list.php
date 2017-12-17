<?php

require __DIR__ . '/../vendor/autoload.php';

// Instantiate the client
$client = new KobaltMusic\Recording\RecordingServiceClient('server:5050', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

// Make the call
list($response, $status) = $client->ListRecordings(new KobaltMusic\Recording\ListRecordingsRequest())->wait();

// Iterate over the results
foreach ($response->getRecordings() as $recording) {
  echo sprintf("ListRecordings() --> '%s' by '%s'\n", $recording->getName(), $recording->getAuthor()->getName());
}

// Profit :)
