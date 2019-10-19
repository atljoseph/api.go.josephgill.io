
THISDIR="$(dirname ${BASH_SOURCE[0]})"
# cd $THISDIR

docker build -t api.josephgill.io:latest .
docker run --name api.josephgill.io --publish 8080:8080 -d --rm api.josephgill.io:latest
