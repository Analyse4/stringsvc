#!/bin/bash
docker build -t gobuild:v0.1 . --no-cache
docker run -p 9001:9001 --rm gobuild:v0.1
