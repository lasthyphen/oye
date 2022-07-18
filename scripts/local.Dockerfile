# syntax=docker/dockerfile:experimental

# This Dockerfile is meant to be used with the build_local_dep_image.sh script
# in order to build an image using the local version of coreth

# Changes to the minimum golang version must also be replicated in
# scripts/ansible/roles/golang_base/defaults/main.yml
# scripts/build_avalanche.sh
# scripts/local.Dockerfile (here)
# Dockerfile
# README.md
# go.mod
FROM golang:1.17.4-buster

RUN mkdir -p /go/src/github.com/lasthyphen

WORKDIR $GOPATH/src/github.com/lasthyphen
COPY avalanchego avalanchego
COPY coreth coreth

WORKDIR $GOPATH/src/github.com/lasthyphen/beacongo
RUN ./scripts/build_avalanche.sh

WORKDIR /go/pkg/mod/github.com/lasthyphen/coreth@v0.8.9
RUN go mod tidy

WORKDIR $GOPATH/src/github.com/lasthyphen/beacongo
RUN ./scripts/build_coreth.sh ../coreth $PWD/build/plugins/evm

RUN ln -sv $GOPATH/src/github.com/lasthyphen/avalanche-byzantine/ /avalanchego
