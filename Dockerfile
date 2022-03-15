FROM golang:alpine

RUN apk add build-base
WORKDIR /app

RUN apk update && apk add bash

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY services/ ./services/
COPY swaggerui/ ./swaggerui/
COPY utils/*.go ./utils/
COPY .openapi-generator-ignore ./
COPY openapi-proxy-service.yaml ./

# maybe there is a better way to use openapi-generator-cli
RUN apk add --update nodejs npm
RUN apk add openjdk11
RUN npm install @openapitools/openapi-generator-cli -g
RUN npx @openapitools/openapi-generator-cli generate -i ./openapi-proxy-service.yaml -g go-server -o ./ --additional-properties=sourceFolder=proxyApi,packageName=proxyApi
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN goimports -w .

RUN go build -o /proxy-service

EXPOSE 8001

CMD [ "/proxy-service" ]