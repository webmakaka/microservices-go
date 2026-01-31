# "Microservices with Go" course project


<br/>

### 04 - Development Environment Setup

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

### 05 - API Gateway and HTTP Servers


```
// POST
$ curl \
    --data '{"userID": "42"}' \
    --header "Content-Type: application/json" \
    --request POST \
    --url http://localhost:8081/trip/preview \
    | jq
```