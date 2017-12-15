FROM php:7.2-cli
RUN apt-get update -y && apt-get install zlib1g-dev
RUN mkdir /usr/src/app
COPY . /usr/src/app
WORKDIR /usr/src/app

RUN pecl install grpc && docker-php-ext-enable grpc
RUN docker-php-ext-install zip
RUN ./composer.phar install
