
THISDIR="$(dirname ${BASH_SOURCE[0]})"
source $THISDIR/docker-variables.sh

docker inspect $IMAGE_NAME:$IMAGE_TAG_FINAL

# check the output