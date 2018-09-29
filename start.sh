#!/bin/bash
docker build -t gobuild:v0.1 . --no-cache
docker run -p 9000:9000 --rm gobuild:v0.1
