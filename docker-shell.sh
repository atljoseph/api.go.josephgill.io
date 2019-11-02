
THISDIR="$(dirname ${BASH_SOURCE[0]})"
source $THISDIR/docker-variables.sh

docker run -it $IMAGE_NAME:$IMAGE_TAG_FINAL --verbose sh
# -v "$IMAGE_VOLUME_LABEL:$IMAGE_VOLUME_DEST"

# opens an interactive shell