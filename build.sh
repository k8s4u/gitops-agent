#!/bin/bash

CGO_ENABLED=0 go build
export ver=v66
docker build . -t k8s4u/gitops-agent:dev-$ver
docker push k8s4u/gitops-agent:dev-$ver

# curl -X POST -H 'Content-Type: application/json' -d '{}' http://127.0.0.1:30778/webhook
