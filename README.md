# "Microservices with Go" course project

<br/>

## 04 - Development Environment Setup

```bash
$ minikube -p marley-minikube addons enable registry
$ eval $(minikube -p marley-minikube docker-env)

$ go mod tidy
$ tilt up
```

<br/>

```
$ go run tools/create_service.go -name trip
```

<br/>

## 05 - API Gateway and HTTP Servers

### 29 - Role of the API Gateway

### 30 - Implementing an HTTP Server

```
// POST
// OK!
$ curl \
    --data '{"userID": "42"}' \
    --header "Content-Type: application/json" \
    --request POST \
    --url http://localhost:8081/trip/preview \
    | jq
```

### 31 - Create a simple HTTP handler

### 32 - External API communication (OSRM)

https://project-osrm.org/

https://project-osrm.org/docs/v5.5.1/api/#route-service

```
// POST
// OK!
$ curl \
    --data '{
        "userID": "12",
            "pickup": {
                "latitude": 37.7960377140795,
                "longitude": -122.42859007644311
            },
            "destination": {
                "latitude": 37.783678887129466,
                "longitude": -122.40777303207217
            }
    }' \
    --header "Content-Type: application/json" \
    --request POST \
    --url http://localhost:8081/trip/preview \
    | jq
```
