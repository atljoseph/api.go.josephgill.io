
# api.josephgill.io

## Run

go run *.go

## Development Database

- [mock-db](./mock-db/ReadMe.md)

## Getting Started

- `go run *.go` will run the `api` locally without docker. If you do this, please also run `docker-compose up mock-db` so that your mock database will run correctly.
- To use Docker locally, run `docker-compose up`, optionally provide the `--build` argument.
- To kill the Docker containers, run `docker-compose down`.
- `sh ./test.sh` will fire api calls at the server.
- If you prefer Postman, then use `postman_collection.json`. 
- NOTE: Still working on getting the UI set up in a docker container / ported over from existing UI. Dev has been pretty heavy on the API side.

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

