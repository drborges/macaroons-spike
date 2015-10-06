build:
	docker build -t authd auth
	docker build -t storaged storage

run: build
	docker run --publish 6060:8080 --name auth-service --detach authd
	docker run --publish 6061:8080 --name storage-service --detach storaged

stop-all:
	docker stop --time=1 auth-service
	docker stop --time=1 storage-service

rm-all: stop-all
	docker rm auth-service
	docker rm storage-service

