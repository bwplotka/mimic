#!/usr/bin/env bash

set +e

# Example YOLO deploy bash script.

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd ${DIR}

echo "Deploying infra-my-mon from `git rev-parse HEAD` to docker compose."

# TODO: Run docker swarm init first.
# TODO: Add secrets.yaml manually.

mkdir -p "/docker-volumes/prometheus-data"
mkdir -p "/docker-volumes/prometheus-config"

# Generate yamls.
go run github.com/bwplotka/gocodeit/examples/infra-my-mon generate --secret-file=secrets.yaml

GEN="gcigen/production/prod-par1-mon0"

cp ${GEN}/deploy/config/prometheus.yaml /docker-volumes/prometheus-config

docker stack -c ${GEN}/deploy/mon-compose.yaml my-mon

