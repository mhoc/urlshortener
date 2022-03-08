#!/bin/bash

# this uses http/httpie, so you'll have to have that to run it. as well as `jq`.
# run `docker compose up` first

DESTINATION="https://hockerman.com";

echo "Generating shortlink for '$DESTINATION'";
SHORT_URL=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink url="$DESTINATION" -b | jq -r '.short_url');

echo "Generated short url: '$SHORT_URL' --> '$DESTINATION'";

echo "resubmitting identical url; it should be the same...";
SHORT_URL_2=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink url="$DESTINATION" -b | jq -r '.short_url');
echo "Generated 2nd short url: '$SHORT_URL_2' --> '$DESTINATION'";

echo "";
echo "Confirming that it works correctly...";
echo "--------";
http $SHORT_URL;
echo "--------";

echo "removing shortlink";
http post localhost:8084/api/URLShortenerV1/RemoveShortlink short_url="$SHORT_URL" -b;
echo "--------";

echo "resubmitting identical url; it should be different now...";
SHORT_URL_3=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink url="$DESTINATION" -b | jq -r '.short_url');
echo "Generated 3nd short url: '$SHORT_URL_3' --> '$DESTINATION'";
echo "--------";

echo "adding expiration";
DESTINATION_2="https://example.com/"
SHORT_URL_4=$(http post localhost:8084/api/URLShortenerV1/CreateShortlink expires_in_seconds=5 url="$DESTINATION_2" -b | jq -r '.short_url');
echo "Generated 4th short url (5s expiration): '$SHORT_URL_4' --> '$DESTINATION_2'";
echo "give it 10 seconds just to be safe...";
sleep 10;
echo "ok lets test it...";
echo "--------";
http $SHORT_URL_4;
echo "--------";