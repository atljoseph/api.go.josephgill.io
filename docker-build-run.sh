
THISDIR="$(dirname ${BASH_SOURCE[0]})"
# cd $THISDIR

docker build -t api.go.josephgill.io:latest .
docker image prune --force
docker run --name api.go.josephgill.io --publish 8080:8080 -d --rm api.go.josephgill.io:latest

sh docker-info.sh
