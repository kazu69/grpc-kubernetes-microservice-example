## Example microservices

> Example: gRPC with golang microservice on Kubernetes

- Defined protocol use gRPC
- Create services with golang
- Create api service with golang
- Build docker image
- Deploy to k8s


```bash
$ make build_greeter_service
$ make build_api

$ make deploy_greeter_service
$ make deploy_api

$ curl http://127.0.0.1:30080/
"Server is running"

$ curl  http://127.0.0.1:30080/greet/YOUR_NAME
"Response is Hello! YOUR_NAME"

$ make delete_greeter_service
$ make delete_api
```
