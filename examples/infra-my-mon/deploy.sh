#!/usr/bin/env bash

# Example YOLO deploy bash script.

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd ${DIR}/../

echo "Deploying infra-my-mon from `git rev-parse HEAD` to docker compose."

# TODO: Run docker swarm init first.
# TODO: Add secrets.yaml manually.

mkdir -p "/docker-volumes/prometheus-data"
mkdir -p "/docker-volumes/prometheus-config"

# Generate yamls.
go run github.com/bwplotka/gocodeit/examples/infra-my-mon generate --secret-file=secrets.yaml

cp gcigen/deploy/config/prometheus.yaml /docker-volumes/prometheus-config

docker stack -c gcigen/deploy/mon-compose.yaml my-mon

