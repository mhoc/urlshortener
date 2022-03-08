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

## Metrics

An example Prometheus + Grafana setup is included, with one metric exported for demonstration
purposes; a counter named `api_shortlink_redirects`, with the label `short_url`, incremented on 
every redirect for each distinct short url. With the stack up, this can be found at
[localhost:3000](http://localhost:3000/explore?orgId=1&left=%5B%22now-1h%22,%22now%22,%22Prometheus%22,%7B%22expr%22:%22rate(api_shortlink_redirects%5B1m%5D)%22%7D,%7B%22mode%22:%22Metrics%22%7D,%7B%22ui%22:%5Btrue,true,true,%22none%22%5D%7D%5D)

- Username: `admin`
- Password: `password`


