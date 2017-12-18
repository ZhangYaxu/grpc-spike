<?php

require __DIR__ . '/../vendor/autoload.php';

if (count($argv) < 2) {
  exit('you have to specify a name for your recording!');
}

// Instantiate the client
$client = new KobaltMusic\Recording\RecordingServiceClient('server:5050', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

// Create the request
$author = new KobaltMusic\Recording\Author();
$author->setName(sprintf('PHP %s', phpversion()));
$recording = new KobaltMusic\Recording\Recording();
$recording->setName($argv[1]);
$recording->setAuthor($author);
$request = new KobaltMusic\Recording\AddRecordingRequest();
$request->setRecording($recording);

// Make the call
$client->addRecording($request)->wait();
echo sprintf("AddRecording() <-- '%s' by '%s'", $recording->getName(), $recording->getAuthor()->getName());

// Profit :)
