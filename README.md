# URL Shortener

This is a URL shortening service, which includes usages of technologies like: Go, Protobuf/Twirp,
Redis, Grafana, and Prometheus.

## Running Locally

The entire stack can be ran with `docker compose up`. This spins up the following services:

| Service           | Port |
|-------------------|------|
| API               | 8084 |
| API Documentation | 8085 |
| Redis             | 6379 |
| Prometheus        | 9090 |
| Grafana           | 3000 |

## API

Auto-generated API documentation can be found in the `docs/` folder, or after running the stack
within a web browser on [localhost:8085](https://localhost:8085/).

The API is unauthenticated. 

For example, to create a new shortlink:

```
curl -X POST localhost:8084/api/URLShortenerV1/CreateShortlink -d '{"url":"https://google.com/"}' -H "Content-Type:application/json"

{"short_url":"http://localhost:8084/FQtXL9Yx"}%
```

To remove that shortlink from the system:

```
curl -X POST localhost:8084/api/URLShortenerV1/RemoveShortlink -d '{"url":"http://localhost:8084/FQtXL9Yx"}' -H "Content-Type:application/json"

{"removed":true}%
```

Protobuf, as an API SDL, was selected as it has a lot of Go-centric codegen tooling available,
whereas other SDLs like GraphQL aren't as mature. However, I've always found gRPC to be difficult
to hack in for smaller projects like this; so Twitch's [Twirp](https://twitchtv.github.io/twirp/) is
a cool alternative which still uses Protobuf, but has at least the one advantage of supporting both
Protocol Buffer & JSON communication formats (depending on the Content-Type specified). 

## Metrics

An example Prometheus + Grafana setup is included, with one metric exported for demonstration
purposes; a counter named `api_shortlink_redirects`, with the label `short_url`, incremented on 
every redirect for each distinct short url. With the stack up, this can be found at
[localhost:3000](http://localhost:3000/explore?orgId=1&left=%5B%22now-1h%22,%22now%22,%22Prometheus%22,%7B%22expr%22:%22rate(api_shortlink_redirects%5B1m%5D)%22%7D,%7B%22mode%22:%22Metrics%22%7D,%7B%22ui%22:%5Btrue,true,true,%22none%22%5D%7D%5D)

- Username: `admin`
- Password: `password`

<img width="782" alt="Screen Shot 2022-03-08 at 10 29 55 AM" src="https://user-images.githubusercontent.com/1148452/157275642-49416a43-e566-4a7f-9dac-cdbb2636a4dc.png">

## Scripts

There's a few scripts included in the `script/` folder which can be used to complete some of the
codegen stuff, along with `./manual_test.sh` which runs through all the API endpoints when the
service is running, outputting the data it returns for manual sanity checking.

## Testing

Automated tests can be ran with:

```
go test ./...
```

## InMemoryStore

In addition to the Redis store, there's also an in-memory store available by starting the api server
without the `REDIS_URL` environment variable configured. Its implementation is in 
`pkg/store/in_memory_store.go`, and it should support everything the redis store does (beside,
of course, persistence).

My thinking is that a mock like this can useful in a test/dev environment, where a full redis 
instance is unavailable. Additionally, it gave more code to cover with tests, and I wanted to 
demonstrate a pattern I enjoy whereby difficult to test components can be abstracted with 
`interface` and mocks.
