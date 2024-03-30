# ratelimit
limit by
- ip
- token


Código HTTP: 429
Mensagem: you have reached the maximum number of requests or actions allowed within a certain time frame

## Build
```shell
docker-compose build
```

## Run
```shell
docker-compose up
```

## Run Test
```shell
# load test by ip
make test-by-ip
```
```shell
# load test with api key
make test-apikey
```

### App Host
```
 http://localhost:8080/

```