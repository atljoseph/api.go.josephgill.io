
THISDIR="$(dirname ${BASH_SOURCE[0]})"
source $THISDIR/docker-variables.sh

docker run -it --entrypoint=./main $IMAGE_NAME:$IMAGE_TAG_FINAL
# -v "$IMAGE_VOLUME_LABEL:$IMAGE_VOLUME_DEST" 

# check the output