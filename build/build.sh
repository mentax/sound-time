#!/bin/sh

ver=v0.1.0

env GOOS=linux   go build -o "sound-time.linux.${ver}"   ../
env GOOS=darwin  go build -o "sound-time.mac.${ver}"     ../
env GOOS=windows go build -o "sound-time.win.${ver}.exe" ../
