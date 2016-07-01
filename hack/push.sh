#!/bin/bash

set -ex

cd $( dirname $0 )/..

TAG=074509403805.dkr.ecr.eu-west-1.amazonaws.com/rebuy-kubernetes-aws-volume-discovery:latest

docker build -t ${TAG} .

docker push ${TAG}
