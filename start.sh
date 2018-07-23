#!/bin/sh

SERVER_PORT=8080
WATCH_SERVER_PORT=8000

open http://localhost:$WATCH_SERVER_PORT
gin -a $SERVER_PORT -p $WATCH_SERVER_PORT run main.go
