#!/bin/bash

sudo yum update -y
sudo yum install git -y
sudo yum install -y golang

echo 'export GOPATH=/home/ec2-user/go' >>~/.bash_profile
echo 'export GOROOT=/usr/lib/golang' >>~/.bash_profile
echo 'export GOTOOLDIR=/usr/lib/golang/pkg/tool/linux_amd64' >>~/.bash_profile
source ~/.bash_profile

go version
go env

go get github.com/wpsmith/fifa-battleplan-matchups/concurrent
cd "$GOPATH/src/github.com/wpsmith/fifa-battleplan-matchups/concurrent"
go build -o c .
./c -n 5
