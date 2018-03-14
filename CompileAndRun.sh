#!/bin/bash

$(go fmt ./...)
$(go build -o main .)

./main