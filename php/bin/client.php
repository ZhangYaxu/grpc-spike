<?php

require __DIR__ . '/../vendor/autoload.php';

$client = new KobaltMusic\Recording\RecordingServiceClient('localhost:5051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

$recordingResponse = $client->ListRecordings(new KobaltMusic\Recording\KBEmpty())->wait();
