.PHONY: protoc
protoc:
	protoc -I./proto --go_out=plugins=grpc:./proto  proto/*.proto

.PHONY: build_greeter_service
build_greeter_service:
	@echo "Build docker image of greeter service ..."
	docker build -t "service/greeter" greeter_service
	@echo "Build finish"

.PHONY: build_api
build_api:
	@echo "Build docker image of api ..."
	docker build -t "service/api" api
	@echo "Build finish"

.PHONY: deploy_greeter_service
deploy_greeter_service:
	@echo "Deploying greeter service in kubernetes ..."
	kubectl apply -f ./deployment/greeter_service.yaml
	@echo "Deploy finish"

.PHONY: deploy_api
deploy_api:
	@echo "Deploying api in kubernetes ..."
	kubectl apply -f ./deployment/api_service.yaml
	@echo "Deploy finish"

.PHONY: delete_greeter_service
delete_greeter_service:
	@echo "Deleting greeter service resource ..."
	kubectl delete svc greeter-service
	@echo "Deleting app server deployment resource ..."
	kubectl delete deployment app-server

.PHONY: delete_api
delete_api:
	@echo "Deleting api service resource ..."
	kubectl delete svc api
	@echo "Deleting api server deployment resource"
	kubectl delete deployment api-server
