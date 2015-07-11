#!/bin/bash -eu

//Installtion of Golang
brew install go hg
mkdir $HOME/go

//Setting gopath 
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:usr/local/opt/go/libexec/bin
mkdir -p $GOPATH/src/github.com/bestupefy/canister
cd go

//Installation of Canister
go get github.com/bestupefy/canister


ORG_PATH="github.com/bestupefy"
REPO_PATH="${ORG_PATH}/canister"

if [ ! -h gopath/src/${REPO_PATH} ]; then
	mkdir -p gopath/src/${ORG_PATH}
	ln -s ../../../.. gopath/src/${REPO_PATH} || exit 255
fi

export GOBIN=${PWD}/bin
export GOPATH=${PWD}/gopath

eval $(go env)

