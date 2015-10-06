# Macaroons Spike

# Running the Spike

First, make sure you have [docker toolbox](https://www.docker.com/toolbox) installed on your machine, then:

```shell
$ make run
```

This will build and run two containers:

1. auth-service: Microservice responsible to issue new macaroon tokens to be used by clients
2. storage-service: Microservice that simulates a file storage container whose single endpoint needs to be validated by a macaroon token

# Getting a Macaroon Token

```shell
$ curl -XPOST http://`docker-machine ip macaroons-vm`:6060/auth/tokens -d '{"id":"borges-id"}' -i
```

Response:

```json
{
  "token": "MDAyOGxvY2F0aW9uIGh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9hdXRoCjAwMWZpZGVudGlmaWVyIGF1dGgtc2VydmljZS1pZAowMDJmc2lnbmF0dXJlIAOHOtv7NpQyyBPz3S1TRpntp9-GGj0W4_PIqdYDv8rZCg=="
}
```

# Using a Macaroon Token On a Request

```shell
$ curl http://`docker-machine ip macaroons-vm`:6061/storage/files -H "X-Auth-Token: MDAyOGxvY2F0aW9uIGh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9hdXRoCjAwMWZpZGVudGlmaWVyIGF1dGgtc2VydmljZS1pZAowMDJmc2lnbmF0dXJlIAOHOtv7NpQyyBPz3S1TRpntp9-GGj0W4_PIqdYDv8rZCg==" -i
```

Providing an invalid macaroon token will fail the request.