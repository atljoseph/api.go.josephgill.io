
# https://www.callicoder.com/docker-golang-image-container-example/

docker build -t api.josephgill.io:latest .
docker run --name api.josephgill.io --publish 8080:8080 -d --rm api.josephgill.io:latest

docker kill api.josephgill.io

