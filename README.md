# proxy-service

swagger urls: \
http://localhost:8001/api/swaggerui/ \

openapi gen:
 ```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/swaggerui/openapi.yaml -g go-server -o /local/ --additional-properties=sourceFolder=proxyApi,packageName=proxyApi
goimports -w .
 ```