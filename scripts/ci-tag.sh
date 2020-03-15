#!/bin/sh
# Login docker
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
# Build Golang Application
go build -o dist/amqp-subscriber
# Build docker image
docker build . -t kainonly/amqp-subscriber:${TRAVIS_TAG}
# Push docker image
docker push kainonly/amqp-subscriber:${TRAVIS_TAG}