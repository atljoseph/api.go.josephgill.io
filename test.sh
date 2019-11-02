
echo '\nGET albums'
curl --location --request GET "http://localhost:8080/v1/api/albums"

echo '\nPOST albums'
curl --location --request POST "http://localhost:8080/v1/api/albums"

echo '\nGET photos'
curl --location --request GET "http://localhost:8080/v1/api/albums/sam-shortline"

