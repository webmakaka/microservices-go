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

<br/>

```
$ go get github.com/gorilla/websocket
```

### 35 - Understanding WebSockets

### 36 - Implementing WebSocket connections

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

```
$ go get google.golang.org/grpc
```

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

<br/>

## 10 - Asynchronous Communication

<br/>

```
$ go get github.com/rabbitmq/amqp091-go
```

<br/>

### 59 - Setting up RabbitMQ

<br/>

### 60 - Connecting to RabbitMQ

<br/>

### 61 - RabbitMQ Services Dependency

<br/>

### 62 - Publishing the First Message

```
$ make generate-proto
```

<br/>

### 63 - Message Durability

<br/>

### 64 - Consuming Messages

<br/>

### 65 - Message Distribution & Acknowledgment

<br/>

### 66 - Fair Dispatch

<br/>

### 67 - The Pub - Sub Pattern

<br/>

### 68 - Implementing the Exchange & Topics flow

<br/>

### 69 - JSON Message Sending & Consuming

<br/>

### 70 - Finding a Suitable Driver

<br/>

## 11 - Real-time Notifications

<br/>

### 71 - Understanding how to notify our users

<br/>

### 72 - WebSocket Connection Manager

<br/>

### 73 - Queue Consumer

<br/>

### 74 - Handling incoming messages from the driver

<br/>

### 75 - Listening for Trip Accept event

<br/>

### 76 - Declining a Trip Request

<br/>

## 12 - Payments

```
$ go get github.com/stripe/stripe-go/v81
```

<br/>

### 77 - Payment Flow Overview

<br/>

### 78 - Payment Service setup

<br/>

### 79 - Adding the Stripe secret key

<br/>

### 80 - Stripe Processor Implementation

<br/>

### 81 - Listening to the Payment Event

<br/>

### 82 - Stripe Payment Webhook

<br/>

## 13 - Observability

```
$ go get go.opentelemetry.io/otel/sdk/trace
$ go get go.opentelemetry.io/otel/exporters/jaeger
$ go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
$ go get go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
```

<br/>

### 83 - Intro to Distributed Tracing

<br/>

### 84 - Setting up Tracing

<br/>

### 85 - Jeager Exporter

<br/>

### 86 - HTTP Instrumentation

<br/>

### 87 - gRPC Instrumentation

<br/>

### 88 - RabbitMQ Instrumentation

<br/>

## 14 - Reliability

<br/>

### 89 - Understanding DLQ and Retries

<br/>

### 90 - Implementing Message Retries
