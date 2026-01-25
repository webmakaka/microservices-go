# "Microservices with Go" course project


### 04 - Development Environment Setup

```bash
$ minikube addons enable registry
$ eval $(minikube -p marley-minikube docker-env)

$ go mod tidy
$ tilt up
```

<br/>

```
$ go run tools/create_service.go -name trip
```