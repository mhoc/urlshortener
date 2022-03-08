#!/bin/bash
protoc \
    --doc_opt=html,index.html \
    --doc_out=./docs \
    ./proto/*.proto