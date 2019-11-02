
THISDIR="$(dirname ${BASH_SOURCE[0]})"
source $THISDIR/docker-variables.sh

# clean containers, images, etc
docker system prune --force

docker kill $IMAGE_NAME

sh docker-info.sh
