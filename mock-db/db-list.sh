
THISDIR="$(dirname ${BASH_SOURCE[0]})"
source $THISDIR/docker-variables.sh

docker exec -i $IMAGE_NAME mysql -uroot -ppassword <<< "show Databases;"

# check the output