build:
	docker build -t authd auth
#	docker build -t storaged storage

run: build
	docker run --publish 6060:8080 --name auth-service --rm authd
#	docker run --publish 6061:8080 --name storage-service --rm storaged