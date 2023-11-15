#!/usr/bin/env bash

git config --global user.email "gg@gg.com"
git config --global user.name "Gigi"
git init
git add .
git commit -m "add all the files"
git tag v0.999

go build -ldflags "-X 'main.GitTag=$(git describe --tags)' -X 'main.Timestamp=$(date -u)'"