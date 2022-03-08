#!/bin/bash

echo "1"
SHORT_URL_1=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink url="https://twitter.com/" -b | jq -r '.short_url');
for i in `seq 1 100`; do
    printf "*";
    http $SHORT_URL_1 > /dev/null;
    sleep 0.5;
done
echo "";

echo "2";
SHORT_URL_2=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink url="https://hockerman.com/" -b | jq -r '.short_url');
for i in `seq 1 200`; do
    printf "*"; 
    http $SHORT_URL_2 > /dev/null; 
    sleep 0.1;
done
echo "";

echo "3";
SHORT_URL_3=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink url="https://example.com/" -b | jq -r '.short_url');
for i in `seq 1 50`; do
    printf "*"; 
    http $SHORT_URL_3 > /dev/null;
    sleep 1;
done
echo "";
