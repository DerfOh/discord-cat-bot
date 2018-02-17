#!/bin/bash

echo "Setting go env for scratch"
export GOOS=linux
export CGO_ENABLED=0

echo "Getting dependencies for discord-cat-bot"
#go get

echo "Build go binary for scratch container"
# go build -a -installsuffix cgo .
go get github.com/ahmetb/govvv
govvv build -o discord-cat-bot .

echo "Get CA Certs from cURL"
wget http://curl.haxx.se/download/curl-7.58.0.tar.bz2
tar xjf curl-7.58.0.tar.bz2

echo "Retrieve ca bundle"
./curl-7.58.0/lib/mk-ca-bundle.pl
echo
echo "Done!"
echo
if [ ! -f config.json ]; then
    echo "Unable to locate config file."
    echo "Make sure that you have placed your config.json in this directory."
    echo "Cleaning up..."
    rm -rf curl-7* certdata.txt ca-bundle.crt discord-cat-bot
    exit
fi

echo "Building docker container"
docker build . -t derfoh/discord-cat-bot

echo "Cleaning up..."
rm -rf curl-7* certdata.txt ca-bundle.crt discord-cat-bot