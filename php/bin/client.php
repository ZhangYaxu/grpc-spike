<?php

require __DIR__ . '/../vendor/autoload.php';

$client = new KobaltMusic\Recording\RecordingServiceClient('localhost:5050', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

$call = $client->ListRecordingsStream(new KobaltMusic\Recording\KBEmpty());
$recordings = $call->responses();

foreach ($recordings as $recording) {
  echo $recording->getName() . "\n";
  echo $recording->getAuthor()->getName() . "\n";
}

list($response, $status) = $client->ListRecordings(new KobaltMusic\Recording\KBEmpty())->wait();

foreach ($response->getRecordings() as $recording) {
  echo $recording->getName() . "\n";
  echo $recording->getAuthor()->getName() . "\n";
}

$author = new KobaltMusic\Recording\Author();
$author->setName('PHP client');
$recording = new KobaltMusic\Recording\Recording();
$recording->setName('Me llevas un anillo');
$recording->setAuthor($author);

$client->addRecording($recording);
