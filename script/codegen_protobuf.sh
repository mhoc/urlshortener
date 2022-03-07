#!/bin/bash
protoc --twirp_out=. --go_out=. proto/urlshortener.proto;