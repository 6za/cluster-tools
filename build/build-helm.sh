#!/bin/bash
shopt -s expand_aliases
source  ~/.bash_profile
export HERE=$PWD
export CHART_DIR=$PWD/docs
export SHA=$(git rev-parse --short HEAD)


docker run -it --rm \
    -v $CHART_DIR:/chart \
    arielev/pybump:1.9.3 \
    bump --file /chart/tunnel-agent/Chart.yaml --level minor


docker run --rm -it \
    -v $CHART_DIR:/chart \
    -w /chart \
    kubebuilder \
    helm package tunnel-agent

docker run --rm -it \
    -w /chart \
    -v $CHART_DIR:/chart \
    kubebuilder \
    helm repo index --url https://6za.github.io/cluster-tools/ --merge index.yaml .
