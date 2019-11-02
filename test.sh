
echo '\ntest albums'
curl --location --request GET "http://localhost:8080/v1/api/albums"

echo '\ntest album photos'
curl --location --request GET "http://localhost:8080/v1/api/album/sam-shortline"
