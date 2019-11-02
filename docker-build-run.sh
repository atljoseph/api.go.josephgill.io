
THISDIR="$(dirname ${BASH_SOURCE[0]})"
source $THISDIR/docker-variables.sh

# cd $THISDIR

docker system prune --force

# remove previous images
docker image rm $IMAGE_NAME

# build a new image
docker build -t $IMAGE_NAME:$IMAGE_TAG_FINAL .

# run the image in a new container
docker run --name $IMAGE_NAME -d --rm $IMAGE_NAME:$IMAGE_TAG_FINAL
# -v "$IMAGE_VOLUME_LABEL:$IMAGE_VOLUME_DEST" 

# clean images, since the first build stage does not clean up its own image
docker system prune --force
# docker image prune --force

# display some info
sh docker-info.sh

