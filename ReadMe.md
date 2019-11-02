
# api.josephgill.io

## Run

go run *.go

## Mock Database

- See `mock-db/ReadMe.md`

## Commands

- `go run *.go` will run the `api` locally without docker.
- `sh ./test.sh` will fire api calls at the server.
- `sh ./docker-build-run.sh` will build and run a docker container with the `api` running inside of it. Do not run the server in docker container and locally at the same time.
- `sh ./docker-kill.sh` will kill the container.

