VM_NAME=macaroons-vm

vm:
	@docker-machine create --driver virtualbox $(VM_NAME) || true
	eval "$(docker-machine env macaroons-vm)"
	docker-machine start $(VM_NAME)
	docker-machine env $(VM_NAME)

build:
	docker build -t authd auth
	docker build -t storaged storage

run: vm build
	docker run --publish 6060:8080 --name auth-service --detach authd
	docker run --publish 6061:8080 --name storage-service --detach storaged

stop-all:
	docker stop --time=1 auth-service
	docker stop --time=1 storage-service

rm-all: stop-all
	docker rm auth-service
	docker rm storage-service

destroy:
	docker-machine rm $(VM_NAME)
