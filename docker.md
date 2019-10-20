
# https://www.callicoder.com/docker-golang-image-container-example/

## Trigger a Multi-stage Build, then cleanup the intermediate build image(s)
docker build -t api.go.josephgill.io:latest .
docker image prune --force

## Run the final output image
docker run --name api.go.josephgill.io --publish 8080:8080 -d --rm api.go.josephgill.io:latest

# Kill the docker container
docker kill api.go.josephgill.io

