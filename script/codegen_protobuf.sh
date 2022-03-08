#!/bin/bash
protoc --twirp_out=. --go_out=. proto/URLShortenerV1.proto;