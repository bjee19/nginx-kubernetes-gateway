#!/usr/bin/env bash

set -e

source "${HOME}"/vars.env

cd nginx-gateway-fabric/tests && make test CI=${CI} TAG="${TAG}" PREFIX="${PREFIX}" NGINX_PREFIX="${NGINX_PREFIX}" NGINX_PLUS_PREFIX="${NGINX_PLUS_PREFIX}" PLUS_ENABLED="${PLUS_ENABLED}" GINKGO_LABEL="${GINKGO_LABEL}" GINKGO_FLAGS="${GINKGO_FLAGS}" PULL_POLICY=Always GW_SERVICE_TYPE=LoadBalancer GW_SVC_GKE_INTERNAL=true NGF_VERSION="${NGF_VERSION}"