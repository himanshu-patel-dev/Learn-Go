# Go and Memcached

## Install Memcached in Docker
```
docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine
```

## Install Postgres in Docker
```
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:12.5-alpine
```
