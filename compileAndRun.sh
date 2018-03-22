#!/bin/bash

$(CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .)

if [ $? != 0 ]; then
    echo "Failed to compile code";
    exit 0;
fi

./main
