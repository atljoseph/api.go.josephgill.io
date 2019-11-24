
# api.josephgill.io

## Run

go run *.go

## Development Database

- [mock-db](./mock-db/ReadMe.md)

## Commands

- `go run *.go` will run the `api` locally without docker.
- To use Docker locally, run `docker-compose up`, optionally provide the `--build` argument.
- To kill the Docker containers, run `docker-compose down`.
- `sh ./test.sh` will fire api calls at the server.

## Packages

- [main](./main.md)
- [apierr](./apierr/ReadMe.md)
- [aws](./aws/ReadMe.md)
- [handlers](./handlers/ReadMe.md)
- [logger](./logger/ReadMe.md)
- [photoDB](./photoDB/ReadMe.md)
- [requester](./requester/ReadMe.md)
- [responder](./responder/ReadMe.md)
- [routes](./routes/ReadMe.md)
- [server](./server/ReadMe.md)

