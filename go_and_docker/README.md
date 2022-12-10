# Building a go app with a specific go version (1.16.15 Here) 

## Building Dockerfile
While in project dir (dir of Dockerfile)
```
docker build -t go_test .
```
Then run container using this
```
docker run --rm --network=my-network --name go_test_cont -p 8000:8000 go_test
```

```curl
curl --location --request POST 'http://localhost:8000/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "First",
    "age": 23,
    "address": {
        "street": "first lane",
        "city": "second city"
    }
}'
```
