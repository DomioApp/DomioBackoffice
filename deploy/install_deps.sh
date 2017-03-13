#!/usr/bin/env bash
set -e

echo Installing Go dependendcies...
cd ~/domiopublic

go get -u -v github.com/fatih/color
go get -u -v github.com/gorilla/mux
go get -u -v github.com/gorilla/websocket
go get -u -v github.com/gorilla/context
go get -u -v github.com/tdewolff/minify
go get -u -v github.com/dgrijalva/jwt-go

echo Go dependencies installed!