
# api.josephgill.io

## Run

go run *.go

## Mock Database

- See `mock-db/ReadMe.md`

## Commands

- `go run *.go` will run the `api` locally without docker.
- To use Docker locally, run `docker-compose up`, optionally provide the `--build` argument.
- To kill the Docker containers, run `docker-compose down`.
- `sh ./test.sh` will fire api calls at the server.


