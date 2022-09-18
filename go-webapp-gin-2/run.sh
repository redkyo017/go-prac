#!/bin/bash

file="app"
if [ -f "$file" ] ; then
    rm "$file"
fi

go build -o app
./app