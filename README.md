# "Microservices with Go"

<br/>

## 04 - Development Environment Setup

```bash
$ export \
    PROFILE=${USER}-minikube

$ minikube --profile ${PROFILE} addons enable registry

$ eval $(minikube --profile ${PROFILE} docker-env)

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

### 33 - Preparing for External API Failures

### 34 - Gracefull shutdown

<br/>

## 06 - WebSockets

### 35 - Understanding WebSockets

### 36 - Implementing WebSocket connections

```
$ go get github.com/gorilla/websocket
```

<br/>

```
$ sudo apt install npm
$ sudo npm install -g wscat
```

<br/>

```
$ wscat -c "ws://localhost:8081/ws/drivers?userID=123&packageSlug=van"
Connected (press CTRL+C to quit)
< {"type":"driver.cmd.register","data":{"id":"123","name":"John Doe","profilePicture":"https://randomuser.me/api/portraits/lego/1.jpg","carPlate":"ABC123","packageSlug":"van"}}
```

<br/>

### 37 - Handling CORS

<br/>

## 07 - Service Communication with gRPC

<br/>

### 38 - gRPC Introduction

https://grpc.io/docs/languages/go/quickstart/

<br/>

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

<br/>

```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

<br/>

https://protobuf.dev/installation/

<br/>

https://github.com/protocolbuffers/protobuf/releases

```
$ sudo apt install -y protobuf-compiler
```

```
$ protoc --version
libprotoc 3.12.4
```

<br/>

### 39 - Protocol Buffer file generation process

<br/>

### 40 - Defining the Trip Proto

```
$ make generate-proto
```

<br/>

### 41 - Implementing the Gateway Trip gRPC Client

```
$ go get google.golang.org/grpc
```

<br/>

### 42 - gRPC Server implementation on Trip Service

<br/>

### 43 - Preview Trip Handler - Part 1

<br/>

### 44 - Preview Trip Handler - Part 2

<br/>

### 45 - Create the Trip Start boilerplate

<br/>

### 46 - Why & What are Ride Fares

<br/>

### 47 - Ride Pricing Estimation

<br/>

### 48 - Implementing the TripStart gRPC handler

```
$ make generate-proto
```

<br/>

## 08 - Kubernetes Essentials

<br/>

## 09 - Drivers Service

```
$ make generate-proto
```
